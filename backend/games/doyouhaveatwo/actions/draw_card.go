package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type DrawCardAction struct {
	Player model.PlayerID
}

func (a *DrawCardAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *DrawCardAction) IsComplete() bool {
	return true
}

func (a *DrawCardAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *DrawCardAction) ToInput() model.Input {
	return &inputs.DrawCardInput{
		Player: a.Player,
	}
}

func (a *DrawCardAction) Type() string {
	return "draw_card"
}

func (a *DrawCardAction) String() string {
	return fmt.Sprintf("draw_card(player=%s)", a.Player)
}
