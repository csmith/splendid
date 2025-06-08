package actions

import (
	"fmt"
	"strings"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCardTargetAnyAction struct {
	Player       model.PlayerID
	CardName     string
	TargetPlayer *model.PlayerID
}

func (a *PlayCardTargetAnyAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardTargetAnyAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayCardTargetAnyAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayCardTargetAnyAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}

	switch a.CardName {
	case "Prince":
		return &inputs.PlayPrinceInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	default:
		return nil
	}
}

func (a *PlayCardTargetAnyAction) Type() string {
	return fmt.Sprintf("play_%s", strings.ToLower(a.CardName))
}

func (a *PlayCardTargetAnyAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_%s(player=%s)", strings.ToLower(a.CardName), a.Player)
	}
	return fmt.Sprintf("play_%s(player=%s, target=%s)", strings.ToLower(a.CardName), a.Player, *a.TargetPlayer)
}

func (a *PlayCardTargetAnyAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player (including self)
	for _, player := range g.Players {
		// Can target self, but not eliminated or protected players
		if !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayCardTargetAnyAction{
				Player:       a.Player,
				CardName:     a.CardName,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
