package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayKingAction struct {
	Player       model.PlayerID
	TargetPlayer *model.PlayerID
}

func (a *PlayKingAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayKingAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayKingAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayKingAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}
	return &inputs.PlayKingInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
	}
}

func (a *PlayKingAction) Type() string {
	return "play_king"
}

func (a *PlayKingAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_king(player=%s)", a.Player)
	}
	return fmt.Sprintf("play_king(player=%s, target=%s)", a.Player, *a.TargetPlayer)
}

func (a *PlayKingAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayKingAction{
				Player:       a.Player,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
