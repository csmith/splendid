package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPrinceAction struct {
	Player       model.PlayerID
	TargetPlayer *model.PlayerID
}

func (a *PlayPrinceAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayPrinceAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayPrinceAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayPrinceAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}
	return &inputs.PlayPrinceInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
	}
}

func (a *PlayPrinceAction) Type() string {
	return "play_prince"
}

func (a *PlayPrinceAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_prince(player=%s)", a.Player)
	}
	return fmt.Sprintf("play_prince(player=%s, target=%s)", a.Player, *a.TargetPlayer)
}

func (a *PlayPrinceAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player (including self)
	for _, player := range g.Players {
		// Can target self, but not eliminated or protected players
		if !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayPrinceAction{
				Player:       a.Player,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
