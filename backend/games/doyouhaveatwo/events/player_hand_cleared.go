package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerHandClearedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerHandClearedEvent) Type() serialization.Specifier {
	return specifier("player_hand_cleared")
}

func (e *PlayerHandClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerHandClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.Hand = []serialization.Redactable[model.Card]{}
	return nil
}
