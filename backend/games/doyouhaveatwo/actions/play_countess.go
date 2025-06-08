package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCountessAction struct {
	Player model.PlayerID
}

func (a *PlayCountessAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCountessAction) IsComplete() bool {
	return true
}

func (a *PlayCountessAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *PlayCountessAction) ToInput() model.Input {
	return &inputs.PlayCountessInput{
		Player: a.Player,
	}
}

func (a *PlayCountessAction) Type() string {
	return "play_countess"
}

func (a *PlayCountessAction) String() string {
	return fmt.Sprintf("play_countess(player=%s)", a.Player)
}
