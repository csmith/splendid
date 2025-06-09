package events

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerActionCompleted model.EventType = "player_action_completed"

type PlayerActionCompletedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerActionCompletedEvent) Type() model.EventType {
	return EventPlayerActionCompleted
}

func (e *PlayerActionCompletedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerActionCompletedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player != nil {
		player.PendingAction = model.Redactable[model.Action]{}
	}
	return nil
}

func (e *PlayerActionCompletedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerActionCompletedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
