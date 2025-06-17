package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type PlayerDiscardPileClearedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerDiscardPileClearedEvent) Type() coremodel.Specifier {
	return specifier("player_discard_pile_cleared")
}

func (e *PlayerDiscardPileClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerDiscardPileClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.DiscardPile = []model.Card{}
	return nil
}
