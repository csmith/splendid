package doyouhaveatwo

import (
	"fmt"
)

type Engine struct {
	Game       Game
	updateChan chan<- GameUpdate
	eventChan  <-chan Event
}

func (e *Engine) applyEvent(event Event) error {
	if err := event.Type.Apply(&event, &e.Game); err != nil {
		return fmt.Errorf("failed to apply event %s: %w", event.Type, err)
	}

	e.updateChan <- GameUpdate{
		Game:             e.Game,
		Event:            event,
		AvailableActions: make(map[PlayerID]Redactable[[]Action]),
	}

	return nil
}

