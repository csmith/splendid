package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerRestored model.EventType = "player_restored"

type PlayerRestoredEvent struct {
	Player model.PlayerID
}

func (e *PlayerRestoredEvent) Type() model.EventType {
	return EventPlayerRestored
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

func (e *PlayerRestoredEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerRestoredEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
