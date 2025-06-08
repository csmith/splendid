package actions

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type StartGameAction struct {
	Player model.PlayerID
}

func (a *StartGameAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *StartGameAction) IsComplete() bool {
	return true
}

func (a *StartGameAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *StartGameAction) ToInput() model.Input {
	return &inputs.StartRoundInput{}
}

func (a *StartGameAction) Type() string {
	return "start_game"
}
