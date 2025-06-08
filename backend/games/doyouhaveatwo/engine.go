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

	if e.actionGenerator == nil {
		return result
	}

	for _, player := range e.Game.Players {
		playerActions := e.actionGenerator.GenerateActionsForPlayer(&e.Game, player.ID)
		if len(playerActions) > 0 {
			result[player.ID] = model.Redactable[[]model.Action]{
				Value:     playerActions,
				VisibleTo: map[model.PlayerID]bool{player.ID: true},
			}
		}
	}

	return result
}

func (e *Engine) ProcessAction(playerID model.PlayerID, action model.Action) error {
	player := e.Game.GetPlayer(playerID)
	if player == nil {
		return fmt.Errorf("player not found: %s", playerID)
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
	if player.PendingAction.Value == nil {
		return e.applyEvent(&events.PlayerActionStartedEvent{
			Player: playerID,
			Action: action,
		})
	} else {
		err := e.applyEvent(&events.PlayerActionUpdatedEvent{
			Player: playerID,
			Action: action,
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
