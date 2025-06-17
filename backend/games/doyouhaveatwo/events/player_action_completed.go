package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerActionCompletedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerActionCompletedEvent) Type() serialization.Specifier {
	return specifier("player_action_completed")
}

func (e *PlayerActionCompletedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerActionCompletedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player != nil {
		player.PendingAction = serialization.Redactable[*serialization.Box[model.Action]]{}
	}
	return nil
}
