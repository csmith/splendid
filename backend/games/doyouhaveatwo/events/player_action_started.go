package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerActionStartedEvent struct {
	Player model.PlayerID                                             `json:"player"`
	Action serialization.Redactable[*serialization.Box[model.Action]] `json:"action"`
}

func (e *PlayerActionStartedEvent) Type() serialization.Specifier {
	return specifier("player_action_started")
}

func (e *PlayerActionStartedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerActionStartedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player != nil {
		player.PendingAction = e.Action
	}
	return nil
}
