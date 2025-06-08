package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerEliminated model.EventType = "player_eliminated"

type PlayerEliminatedEvent struct {
	Player model.PlayerID
}

func (e *PlayerEliminatedEvent) Type() model.EventType {
	return EventPlayerEliminated
}

func (e *PlayerEliminatedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerEliminatedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsOut = true
	return nil
}
