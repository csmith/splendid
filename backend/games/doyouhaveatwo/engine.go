package doyouhaveatwo

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type Engine struct {
	Game            model.Game
	EventHistory    []model.Event
	updateChan      chan<- model.GameUpdate
	actionGenerator actions.ActionGenerator
	eventLogger     EventLogger
}

func NewEngine(updateChan chan<- model.GameUpdate, eventLogger EventLogger) *Engine {
	return &Engine{
		Game: model.Game{
			Players:       []*model.Player{},
			Deck:          []model.Redactable[model.Card]{},
			CurrentPlayer: 0,
			Round:         0,
			Phase:         model.PhaseSetup,
			TokensToWin:   4,
		},
		EventHistory:    []model.Event{},
		updateChan:      updateChan,
		actionGenerator: &actions.DefaultActionGenerator{},
		eventLogger:     eventLogger,
	}
}

func (e *Engine) processInput(input model.Input) error {
	var applyError error

	// Create callback function that applies events immediately
	apply := func(event model.Event) {
		if applyError != nil {
			return // Skip if we already have an error
		}
		applyError = e.applyEvent(event)
	}

	// Apply the input with the callback
	err := input.Apply(&e.Game, apply)
	if err != nil {
		return fmt.Errorf("failed to process input %s: %w", input.Type(), err)
	}

	// Check for any errors that occurred during event application
	if applyError != nil {
		return fmt.Errorf("failed to apply event during input %s: %w", input.Type(), applyError)
	}

	return nil
}

func (e *Engine) applyEvent(event model.Event) error {
	if err := event.Apply(&e.Game); err != nil {
		return fmt.Errorf("failed to apply event %s: %w", event.Type(), err)
	}

	// Store event in history
	e.EventHistory = append(e.EventHistory, event)

	// Log event to persistent storage
	if e.eventLogger != nil {
		if err := e.eventLogger.LogEvent(event); err != nil {
			return fmt.Errorf("failed to log event %s: %w", event.Type(), err)
		}
	}

	e.updateChan <- model.GameUpdate{
		Game:             e.Game,
		Event:            &serialization.Box[model.Event]{Value: event},
		AvailableActions: e.generateAvailableActions(),
	}

	return nil
}

func (e *Engine) generateAvailableActions() map[model.PlayerID]model.Redactable[[]serialization.Box[model.Action]] {
	result := make(map[model.PlayerID]model.Redactable[[]serialization.Box[model.Action]])

	for _, player := range e.Game.Players {
		playerActions := e.actionGenerator.GenerateActionsForPlayer(&e.Game, player.ID)

		if len(playerActions) > 0 {
			// Convert to boxes
			boxedActions := make([]serialization.Box[model.Action], len(playerActions))
			for i, action := range playerActions {
				boxedActions[i] = serialization.Box[model.Action]{Value: action}
			}
			result[player.ID] = model.NewRedactable(boxedActions, player.ID)
		}
	}

	return result
}

func (e *Engine) validateAction(playerID model.PlayerID, action model.Action) error {
	// If no available actions have been generated yet, generate them now
	availableForPlayer := e.actionGenerator.GenerateActionsForPlayer(&e.Game, playerID)

	// Check if the submitted action matches any of the available actions
	for _, availableAction := range availableForPlayer {
		if e.actionsMatch(action, availableAction) {
			return nil
		}
	}

	return fmt.Errorf("action %s is not available for player %s, only: %s", action, playerID, availableForPlayer)
}

func (e *Engine) actionsMatch(submitted, available model.Action) bool {
	return submitted.String() == available.String()
}

func (e *Engine) ProcessAction(playerID model.PlayerID, action model.Action) error {
	player := e.Game.GetPlayer(playerID)
	if player == nil {
		return fmt.Errorf("player not found: %s", playerID)
	}

	// Validate that the action was previously generated for this player
	if err := e.validateAction(playerID, action); err != nil {
		return err
	}

	// If action is already complete, execute immediately
	if action.IsComplete() {
		concreteInput := action.ToInput()
		if concreteInput != nil {
			err := e.processInput(concreteInput)
			if err != nil {
				return err
			}
		}

		// Clear the pending action after processing
		return e.applyEvent(&events.PlayerActionCompletedEvent{
			Player: playerID,
		})
	}

	// If no pending action, start it; otherwise update it
	if player.PendingAction.Value() == nil {
		return e.applyEvent(&events.PlayerActionStartedEvent{
			Player: playerID,
			Action: model.NewRedactable(&serialization.Box[model.Action]{Value: action}, playerID),
		})
	} else {
		// Check if the action is now complete after update
		if action.IsComplete() {
			// Execute the concrete input first
			concreteInput := action.ToInput()
			if concreteInput != nil {
				err := e.processInput(concreteInput)
				if err != nil {
					return err
				}
			}

			// Clear the pending action after processing
			return e.applyEvent(&events.PlayerActionCompletedEvent{
				Player: playerID,
			})
		} else {
			return e.applyEvent(&events.PlayerActionUpdatedEvent{
				Player: playerID,
				Action: model.NewRedactable(&serialization.Box[model.Action]{Value: action}, playerID),
			})
		}
	}
}

func (e *Engine) ProcessServerAction(action model.Action) error {
	// Execute the concrete input directly without validation or pending action handling
	concreteInput := action.ToInput()
	if concreteInput != nil {
		return e.processInput(concreteInput)
	}
	return nil
}
