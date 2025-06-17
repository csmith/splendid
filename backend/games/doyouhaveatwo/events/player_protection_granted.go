package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type PlayerProtectionGrantedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerProtectionGrantedEvent) Type() coremodel.Specifier {
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
