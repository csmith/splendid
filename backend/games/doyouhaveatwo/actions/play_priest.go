package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPriestAction struct {
	Player       model.PlayerID
	TargetPlayer *model.PlayerID
}

func (a *PlayPriestAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayPriestAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayPriestAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayPriestAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}
	return &inputs.PlayPriestInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
	}
}

func (a *PlayPriestAction) Type() string {
	return "play_priest"
}

func (a *PlayPriestAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_priest(player=%s)", a.Player)
	}
	return fmt.Sprintf("play_priest(player=%s, target=%s)", a.Player, *a.TargetPlayer)
}

func (a *PlayPriestAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayPriestAction{
				Player:       a.Player,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
