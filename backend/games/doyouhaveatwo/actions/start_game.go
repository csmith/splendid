package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type StartGameAction struct {
	Player model.PlayerID `json:"player"`
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

func (a *StartGameAction) Type() serialization.Specifier {
	return specifier("start_game")
}

func (a *StartGameAction) String() string {
	return fmt.Sprintf("start_game(player=%s)", a.Player)
}
