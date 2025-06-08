package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCardNoTargetAction struct {
	Player model.PlayerID
	Card   model.Card
}

func (a *PlayCardNoTargetAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardNoTargetAction) IsComplete() bool {
	return true
}

func (a *PlayCardNoTargetAction) NextActions(g *model.Game) []model.Action {
	return nil
}

func (a *PlayCardNoTargetAction) ToInput() model.Input {
	switch a.Card {
	case model.CardHandmaid:
		return &inputs.PlayHandmaidInput{
			Player: a.Player,
		}
	case model.CardCountess:
		return &inputs.PlayCountessInput{
			Player: a.Player,
		}
	case model.CardPrincess:
		return &inputs.PlayPrincessInput{
			Player: a.Player,
		}
	default:
		return nil
	}
}

func (a *PlayCardNoTargetAction) Type() string {
	return fmt.Sprintf("play_%s", a.Card.Name())
}

func (a *PlayCardNoTargetAction) String() string {
	return fmt.Sprintf("play_%s(player=%s)", a.Card.Name(), a.Player)
}
