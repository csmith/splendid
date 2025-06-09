package events

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerActionUpdated model.EventType = "player_action_updated"

type PlayerActionUpdatedEvent struct {
	Player model.PlayerID                 `json:"player"`
	Action model.Redactable[model.Action] `json:"action"`
}

func (e *PlayerActionUpdatedEvent) Type() model.EventType {
	return EventPlayerActionUpdated
}

func (e *PlayerActionUpdatedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerActionUpdatedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player != nil {
		player.PendingAction = e.Action
	}
	return nil
}

func (e *PlayerActionUpdatedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerActionUpdatedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
