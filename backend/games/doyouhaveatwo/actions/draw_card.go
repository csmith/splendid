package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type DrawCardAction struct {
	Player model.PlayerID `json:"player"`
}

func (a *DrawCardAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *DrawCardAction) IsComplete() bool {
	return true
}

func (a *DrawCardAction) NextActions(g *model.Game) []model.GameAction {
	return nil
}

func (a *DrawCardAction) ToInput() model.Input {
	return &inputs.DrawCardInput{
		Player: a.Player,
	}
}

func (a *DrawCardAction) Type() coremodel.Specifier {
	return specifier("draw_card")
}

func (a *DrawCardAction) String() string {
	return fmt.Sprintf("draw_card(player=%s)", a.Player)
}
