package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type WinnerDeclaredEvent struct {
	Winner model.PlayerID `json:"winner"`
}

func (e *WinnerDeclaredEvent) Type() coremodel.Specifier {
	return specifier("winner_declared")
}

func (e *WinnerDeclaredEvent) PlayerID() *model.PlayerID {
	return &e.Winner
}

func (e *WinnerDeclaredEvent) Apply(g *model.Game) error {
	// This event is primarily for client notification
	// The game state doesn't need to be modified beyond phase changes
	return nil
}
