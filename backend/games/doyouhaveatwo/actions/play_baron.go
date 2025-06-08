package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayBaronAction struct {
	Player       model.PlayerID
	TargetPlayer *model.PlayerID
}

func (a *PlayBaronAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayBaronAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayBaronAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayBaronAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}
	return &inputs.PlayBaronInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
	}
}

func (a *PlayBaronAction) Type() string {
	return "play_baron"
}

func (a *PlayBaronAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_baron(player=%s)", a.Player)
	}
	return fmt.Sprintf("play_baron(player=%s, target=%s)", a.Player, *a.TargetPlayer)
}

func (a *PlayBaronAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayBaronAction{
				Player:       a.Player,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
