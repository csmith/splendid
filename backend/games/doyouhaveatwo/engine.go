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
	events, err := input.Apply(&e.Game)
	if err != nil {
		return fmt.Errorf("failed to process input %s: %w", input.Type(), err)
	}

	for _, event := range events {
		if err := e.applyEvent(event); err != nil {
			return err
		}
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
