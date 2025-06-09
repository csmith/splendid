package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventPlayerProtectionCleared model.EventType = "player_protection_cleared"

type PlayerProtectionClearedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerProtectionClearedEvent) Type() serialization.Specifier {
	return specifier("player_protection_cleared")
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
