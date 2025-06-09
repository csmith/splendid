package actions

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type DiscardCardAction struct {
	Player   model.PlayerID `json:"player"`
	CardName string         `json:"card_name"`
}

func (a *DiscardCardAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *DiscardCardAction) IsComplete() bool {
	return true
}

func (a *DiscardCardAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *DiscardCardAction) ToInput() model.Input {
	return &inputs.DiscardCardInput{
		Player:   a.Player,
		CardName: a.CardName,
	}
}

func (a *DiscardCardAction) Type() string {
	return "discard_card"
}

func (a *DiscardCardAction) String() string {
	return fmt.Sprintf("discard_card(player=%s, card=%s)", a.Player, a.CardName)
}

func (a *DiscardCardAction) MarshalJSON() ([]byte, error) {
	type Alias DiscardCardAction
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  a.Type(),
		Alias: (*Alias)(a),
	})
}
