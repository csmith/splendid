package doyouhaveatwo

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Engine struct {
	Game       model.Game
	updateChan chan<- model.GameUpdate
	eventChan  <-chan model.Event
}

func (e *Engine) applyEvent(event model.Event) error {
	if err := event.Apply(&e.Game); err != nil {
		return fmt.Errorf("failed to apply event %s: %w", event.Type(), err)
	}

	e.updateChan <- model.GameUpdate{
		Game:             e.Game,
		Event:            event,
		AvailableActions: make(map[model.PlayerID]model.Redactable[[]model.Action]),
	}

	return nil
}
