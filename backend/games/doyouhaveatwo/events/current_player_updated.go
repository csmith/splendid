package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventCurrentPlayerUpdated model.EventType = "current_player_updated"

type CurrentPlayerUpdatedEvent struct {
	NewCurrentPlayer int `json:"current_player"`
}

func (e *CurrentPlayerUpdatedEvent) Type() serialization.Specifier {
	return specifier("current_player_updated")
}

func (e *CurrentPlayerUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *CurrentPlayerUpdatedEvent) Apply(g *model.Game) error {
	g.CurrentPlayer = e.NewCurrentPlayer
	return nil
}
