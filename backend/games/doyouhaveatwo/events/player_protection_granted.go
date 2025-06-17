package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerProtectionGrantedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerProtectionGrantedEvent) Type() serialization.Specifier {
	return specifier("player_protection_granted")
}

func (e *PlayerProtectionGrantedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerProtectionGrantedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsProtected = true
	return nil
}
