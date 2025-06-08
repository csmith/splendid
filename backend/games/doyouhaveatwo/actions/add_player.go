package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type AddPlayerAction struct {
	NewPlayerID   model.PlayerID
	NewPlayerName string
}

func (a *AddPlayerAction) PlayerID() model.PlayerID {
	return a.NewPlayerID
}

func (a *AddPlayerAction) IsComplete() bool {
	return a.NewPlayerName != ""
}

func (a *AddPlayerAction) NextActions(g *model.Game) []model.Action {
	// This is a simple action that doesn't require multiple steps
	return nil
}

func (a *AddPlayerAction) ToInput() model.Input {
	return &inputs.AddPlayerInput{
		NewPlayerID:   a.NewPlayerID,
		NewPlayerName: a.NewPlayerName,
	}
}

func (a *AddPlayerAction) Type() string {
	return "add_player"
}

func (a *AddPlayerAction) String() string {
	if a.NewPlayerName == "" {
		return fmt.Sprintf("add_player(id=%s)", a.NewPlayerID)
	}
	return fmt.Sprintf("add_player(id=%s, name=%s)", a.NewPlayerID, a.NewPlayerName)
}
