package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type PlayerRestoredEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerRestoredEvent) Type() coremodel.Specifier {
	return specifier("player_restored")
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
