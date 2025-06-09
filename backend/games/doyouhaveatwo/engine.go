package doyouhaveatwo

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Engine struct {
	Game            model.Game
	EventHistory    []model.Event
	updateChan      chan<- model.GameUpdate
	actionGenerator actions.ActionGenerator
}

func NewEngine(updateChan chan<- model.GameUpdate) *Engine {
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

	e.updateChan <- model.GameUpdate{
		Game:             e.Game,
		Event:            event,
		AvailableActions: e.generateAvailableActions(),
	}

	return nil
}

func (e *Engine) generateAvailableActions() map[model.PlayerID]model.Redactable[[]model.Action] {
	result := make(map[model.PlayerID]model.Redactable[[]model.Action])

	for _, player := range e.Game.Players {
		playerActions := e.actionGenerator.GenerateActionsForPlayer(&e.Game, player.ID)

		if len(playerActions) > 0 {
			result[player.ID] = model.NewRedactable(playerActions, player.ID)
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
			return e.processInput(concreteInput)
		}
		return nil
	}

	// If no pending action, start it; otherwise update it
	if player.PendingAction.Value() == nil {
		return e.applyEvent(&events.PlayerActionStartedEvent{
			Player: playerID,
			Action: model.NewRedactable(action, playerID),
		})
	} else {
		err := e.applyEvent(&events.PlayerActionUpdatedEvent{
			Player: playerID,
			Action: model.NewRedactable(action, playerID),
		})
		if err != nil {
			return err
		}

		// Check if the action is now complete after update
		if action.IsComplete() {
			// Clear the pending action
			err := e.applyEvent(&events.PlayerActionCompletedEvent{
				Player: playerID,
			})
			if err != nil {
				return err
			}

			// Execute the concrete input
			concreteInput := action.ToInput()
			if concreteInput != nil {
				return e.processInput(concreteInput)
			}
		}
	}

	return nil
}

func (e *Engine) ProcessServerAction(action model.Action) error {
	// Execute the concrete input directly without validation or pending action handling
	concreteInput := action.ToInput()
	if concreteInput != nil {
		return e.processInput(concreteInput)
	}
	return nil
}
