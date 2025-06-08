package events

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

const EventCurrentPlayerUpdated model.EventType = "current_player_updated"

type CurrentPlayerUpdatedEvent struct {
	NewCurrentPlayer int
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
