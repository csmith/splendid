package events

import (
	"encoding/json"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventWinnerDeclared model.EventType = "winner_declared"

type WinnerDeclaredEvent struct {
	Winner model.PlayerID
}

func (e *WinnerDeclaredEvent) Type() model.EventType {
	return EventWinnerDeclared
}

func (e *WinnerDeclaredEvent) PlayerID() *model.PlayerID {
	return &e.Winner
}

func (e *WinnerDeclaredEvent) Apply(g *model.Game) error {
	// This event is primarily for client notification
	// The game state doesn't need to be modified beyond phase changes
	return nil
}

func (e *WinnerDeclaredEvent) MarshalJSON() ([]byte, error) {
	type Alias WinnerDeclaredEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
