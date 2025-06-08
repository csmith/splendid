package doyouhaveatwo

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Engine struct {
	Game         model.Game
	EventHistory []model.Event
	updateChan   chan<- model.GameUpdate
	inputChan    <-chan inputs.Input
}

func (e *Engine) processInput(input inputs.Input) error {
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
		AvailableActions: make(map[model.PlayerID]model.Redactable[[]model.Action]),
	}

	return nil
}
