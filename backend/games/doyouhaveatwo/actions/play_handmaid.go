package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayHandmaidAction struct {
	Player model.PlayerID
}

func (a *PlayHandmaidAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayHandmaidAction) IsComplete() bool {
	return true
}

func (a *PlayHandmaidAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *PlayHandmaidAction) ToInput() model.Input {
	return &inputs.PlayHandmaidInput{
		Player: a.Player,
	}
}

func (a *PlayHandmaidAction) Type() string {
	return "play_handmaid"
}

func (a *PlayHandmaidAction) String() string {
	return fmt.Sprintf("play_handmaid(player=%s)", a.Player)
}
