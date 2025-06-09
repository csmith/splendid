package events

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventCurrentPlayerUpdated model.EventType = "current_player_updated"

type CurrentPlayerUpdatedEvent struct {
	NewCurrentPlayer int `json:"current_player"`
}

func (e *CurrentPlayerUpdatedEvent) Type() model.EventType {
	return EventCurrentPlayerUpdated
}

func (e *CurrentPlayerUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *CurrentPlayerUpdatedEvent) Apply(g *model.Game) error {
	g.CurrentPlayer = e.NewCurrentPlayer
	return nil
}

func (e *CurrentPlayerUpdatedEvent) MarshalJSON() ([]byte, error) {
	type Alias CurrentPlayerUpdatedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
