package events

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

const EventRoundUpdated model.EventType = "round_updated"

type RoundUpdatedEvent struct {
	NewRound int
}

func (e *RoundUpdatedEvent) Type() model.EventType {
	return EventRoundUpdated
}

func (e *RoundUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *RoundUpdatedEvent) Apply(g *model.Game) error {
	g.Round = e.NewRound
	return nil
}
