package actions

import (
	"fmt"
	"strings"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type PlayCardTargetOthersAction struct {
	Player       model.PlayerID  `json:"player"`
	CardName     string          `json:"card_name"`
	TargetPlayer *model.PlayerID `json:"target_player"`
}

func (a *PlayCardTargetOthersAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardTargetOthersAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayCardTargetOthersAction) NextActions(g *model.Game) []model.GameAction {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayCardTargetOthersAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}

	switch a.CardName {
	case "Baron":
		return &inputs.PlayBaronInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	case "Priest":
		return &inputs.PlayPriestInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	case "King":
		return &inputs.PlayKingInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	default:
		return nil
	}
}

func (a *PlayCardTargetOthersAction) Type() coremodel.Specifier {
	return specifier("play_card_target_others")
}

func (a *PlayCardTargetOthersAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_%s(player=%s)", strings.ToLower(a.CardName), a.Player)
	}
	return fmt.Sprintf("play_%s(player=%s, target=%s)", strings.ToLower(a.CardName), a.Player, *a.TargetPlayer)
}

func (a *PlayCardTargetOthersAction) generateTargetActions(g *model.Game) []model.GameAction {
	var actions []model.GameAction

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayCardTargetOthersAction{
				Player:       a.Player,
				CardName:     a.CardName,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
