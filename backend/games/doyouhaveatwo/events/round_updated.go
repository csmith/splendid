package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type RoundUpdatedEvent struct {
	NewRound int `json:"round"`
}

func (e *RoundUpdatedEvent) Type() coremodel.Specifier {
	return specifier("round_updated")
}

func (e *RoundUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *RoundUpdatedEvent) Apply(g *model.Game) error {
	g.Round = e.NewRound
	return nil
}
