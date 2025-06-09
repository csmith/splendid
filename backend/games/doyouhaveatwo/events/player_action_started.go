package events

import (
	"encoding/json"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerActionStarted model.EventType = "player_action_started"

type PlayerActionStartedEvent struct {
	Player model.PlayerID
	Action model.Action
}

func (e *PlayerActionStartedEvent) Type() model.EventType {
	return EventPlayerActionStarted
}

func (e *PlayerActionStartedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerActionStartedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player != nil {
		player.PendingAction = model.Redactable[model.Action]{
			Value:     e.Action,
			VisibleTo: map[model.PlayerID]bool{e.Player: true},
		}
	}
	return nil
}

func (e *PlayerActionStartedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerActionStartedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
