package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPrincessAction struct {
	Player model.PlayerID
}

func (a *PlayPrincessAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayPrincessAction) IsComplete() bool {
	return true
}

func (a *PlayPrincessAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *PlayPrincessAction) ToInput() model.Input {
	return &inputs.PlayPrincessInput{
		Player: a.Player,
	}
}

func (a *PlayPrincessAction) Type() string {
	return "play_princess"
}

func (a *PlayPrincessAction) String() string {
	return fmt.Sprintf("play_princess(player=%s)", a.Player)
}
