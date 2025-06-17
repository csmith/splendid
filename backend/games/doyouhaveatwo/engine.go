package doyouhaveatwo

import (
	"fmt"

	coremodel "github.com/csmith/splendid/backend/model"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Engine struct {
	Game            model.Game
	EventHistory    []model.Event
	PendingActions  map[model.PlayerID]model.GameAction
	updateChan      chan<- model.GameUpdate
	actionGenerator coremodel.ActionGenerator[model.Game]
	eventLogger     EventLogger
}

func NewEngine(game model.Game, updateChan chan<- model.GameUpdate, eventLogger EventLogger) *Engine {
	return &Engine{
		Game:            game,
		EventHistory:    []model.Event{},
		PendingActions:  make(map[model.PlayerID]model.GameAction),
		updateChan:      updateChan,
		actionGenerator: &actions.Generator{},
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
		return fmt.Errorf("failed to process input: %w", err)
	}

	// Check for any errors that occurred during event application
	if applyError != nil {
		return fmt.Errorf("failed to apply event during input processing: %w", applyError)
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
		Game:             &e.Game,
		Event:            &coremodel.Box[model.Event]{Value: event},
		AvailableActions: e.generateAvailableActions(),
	}

	return nil
}

func (e *Engine) generateAvailableActions() map[model.PlayerID]coremodel.Redactable[[]*coremodel.Box[model.GameAction]] {
	result := make(map[model.PlayerID]coremodel.Redactable[[]*coremodel.Box[model.GameAction]])

	for _, player := range e.Game.Players {
		var playerActions []model.GameAction

		// If player has a pending action, return next steps
		if pendingAction, exists := e.PendingActions[player.ID]; exists && pendingAction != nil {
			playerActions = pendingAction.NextActions(&e.Game)
		} else {
			// Otherwise, generate initial actions based on game state
			playerActions = e.actionGenerator.GenerateActionsForPlayer(&e.Game, player.ID)
		}

		if len(playerActions) > 0 {
			// Convert to boxes
			boxedActions := make([]*coremodel.Box[model.GameAction], len(playerActions))
			for i, action := range playerActions {
				boxedActions[i] = &coremodel.Box[model.GameAction]{Value: action}
			}
			result[player.ID] = coremodel.NewRedactable(boxedActions, player.ID)
		}
	}

	return result
}

func (e *Engine) validateAction(playerID model.PlayerID, action model.GameAction) error {
	// If no available actions have been generated yet, generate them now
	var availableForPlayer []model.GameAction
	if pendingAction, exists := e.PendingActions[playerID]; exists && pendingAction != nil {
		availableForPlayer = pendingAction.NextActions(&e.Game)
	} else {
		availableForPlayer = e.actionGenerator.GenerateActionsForPlayer(&e.Game, playerID)
	}

	// Check if the submitted action matches any of the available actions
	for _, availableAction := range availableForPlayer {
		if e.actionsMatch(action, availableAction) {
			return nil
		}
	}

	return fmt.Errorf("action %s is not available for player %s, only: %s", action, playerID, availableForPlayer)
}

func (e *Engine) actionsMatch(submitted, available model.GameAction) bool {
	return submitted.String() == available.String()
}

func (e *Engine) ProcessAction(playerID model.PlayerID, action model.GameAction) error {
	player := e.Game.GetPlayer(playerID)
	if player == nil {
		return fmt.Errorf("player not found: %s", playerID)
	}

	// Validate that the action was previously generated for this player
	if err := e.validateAction(playerID, action); err != nil {
		return err
	}

	// Update pending action state
	e.PendingActions[playerID] = action

	// If action is complete, execute it
	if action.IsComplete() {
		concreteInput := action.ToInput()
		if concreteInput != nil {
			if err := e.processInput(concreteInput); err != nil {
				return err
			}
		}
		// Clear the pending action after processing
		delete(e.PendingActions, playerID)
	} else {
		// Send update for incomplete action (complete actions send updates via processInput)
		e.updateChan <- model.GameUpdate{
			Game:             &e.Game,
			Event:            nil,
			AvailableActions: e.generateAvailableActions(),
		}
	}

	return nil
}

func (e *Engine) ProcessServerAction(action model.GameAction) error {
	// Execute the concrete input directly without validation or pending action handling
	concreteInput := action.ToInput()
	if concreteInput != nil {
		return e.processInput(concreteInput)
	}
	return nil
}
