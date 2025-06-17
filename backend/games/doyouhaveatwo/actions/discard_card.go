package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
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

func (a *DiscardCardAction) NextActions(g *model.Game) []model.GameAction {
	return nil
}

func (a *DiscardCardAction) ToInput() model.Input {
	return &inputs.DiscardCardInput{
		Player:   a.Player,
		CardName: a.CardName,
	}
}

func (a *DiscardCardAction) Type() coremodel.Specifier {
	return specifier("discard_card")
}

func (a *DiscardCardAction) String() string {
	return fmt.Sprintf("discard_card(player=%s, card=%s)", a.Player, a.CardName)
}
