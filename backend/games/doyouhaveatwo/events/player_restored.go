package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerRestored model.EventType = "player_restored"

type PlayerRestoredEvent struct {
	Player model.PlayerID
}

func (e *PlayerRestoredEvent) Type() model.EventType {
	return EventPlayerRestored
}

func (e *PlayerRestoredEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerRestoredEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsOut = false
	return nil
}
