package events

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

const EventPlayerActionCompleted model.EventType = "player_action_completed"

type PlayerActionCompletedEvent struct {
	Player model.PlayerID
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
