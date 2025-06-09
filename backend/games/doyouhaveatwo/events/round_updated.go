package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventRoundUpdated model.EventType = "round_updated"

type RoundUpdatedEvent struct {
	NewRound int `json:"round"`
}

func (e *RoundUpdatedEvent) Type() serialization.Specifier {
	return specifier("round_updated")
}

func (e *RoundUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *RoundUpdatedEvent) Apply(g *model.Game) error {
	g.Round = e.NewRound
	return nil
}
