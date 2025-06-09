package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventPlayerActionUpdated model.EventType = "player_action_updated"

type PlayerActionUpdatedEvent struct {
	Player model.PlayerID                                             `json:"player"`
	Action serialization.Redactable[*serialization.Box[model.Action]] `json:"action"`
}

func (e *PlayerActionUpdatedEvent) Type() serialization.Specifier {
	return specifier("player_action_updated")
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
