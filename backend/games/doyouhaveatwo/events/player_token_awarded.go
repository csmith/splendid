package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerTokenAwardedEvent struct {
	Player model.PlayerID `json:"player"`
	Tokens int            `json:"tokens"`
}

func (e *PlayerTokenAwardedEvent) Type() serialization.Specifier {
	return specifier("player_token_awarded")
}

func (e *PlayerTokenAwardedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerTokenAwardedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.TokenCount += e.Tokens
	return nil
}
