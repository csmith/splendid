package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerProtectionCleared model.EventType = "player_protection_cleared"

type PlayerProtectionClearedEvent struct {
	Player model.PlayerID
}

func (e *PlayerProtectionClearedEvent) Type() model.EventType {
	return EventPlayerProtectionCleared
}

func (e *PlayerProtectionClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerProtectionClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsProtected = false
	return nil
}
