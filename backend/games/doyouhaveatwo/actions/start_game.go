package actions

import (
	"encoding/json"
	"fmt"

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

func (a *StartGameAction) String() string {
	return fmt.Sprintf("start_game(player=%s)", a.Player)
}

func (a *StartGameAction) MarshalJSON() ([]byte, error) {
	type Alias StartGameAction
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  a.Type(),
		Alias: (*Alias)(a),
	})
}
