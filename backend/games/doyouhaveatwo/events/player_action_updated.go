package events

import (
	"encoding/json"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerActionUpdated model.EventType = "player_action_updated"

type PlayerActionUpdatedEvent struct {
	Player model.PlayerID
	Action model.Action
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
		player.PendingAction = model.Redactable[model.Action]{
			Value:     e.Action,
			VisibleTo: map[model.PlayerID]bool{e.Player: true},
		}
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
