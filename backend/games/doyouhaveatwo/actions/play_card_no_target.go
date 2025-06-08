package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCardNoTargetAction struct {
	Player   model.PlayerID
	CardName string
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
	switch a.CardName {
	case "Handmaid":
		return &inputs.PlayHandmaidInput{
			Player: a.Player,
		}
	case "Countess":
		return &inputs.PlayCountessInput{
			Player: a.Player,
		}
	case "Princess":
		return &inputs.PlayPrincessInput{
			Player: a.Player,
		}
	default:
		return nil
	}
}

func (a *PlayCardNoTargetAction) Type() string {
	return fmt.Sprintf("play_%s", a.CardName)
}

func (a *PlayCardNoTargetAction) String() string {
	return fmt.Sprintf("play_%s(player=%s)", a.CardName, a.Player)
}
