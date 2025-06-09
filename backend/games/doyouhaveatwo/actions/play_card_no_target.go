package actions

import (
	"fmt"
	"strings"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayCardNoTargetAction struct {
	Player   model.PlayerID `json:"player"`
	CardName string         `json:"card_name"`
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

func (a *PlayCardNoTargetAction) Type() serialization.Specifier {
	return specifier("play_card_no_target")
}

func (a *PlayCardNoTargetAction) String() string {
	return fmt.Sprintf("play_%s(player=%s)", strings.ToLower(a.CardName), a.Player)
}
