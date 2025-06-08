package actions

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCardTargetOthersAction struct {
	Player       model.PlayerID
	Card         model.Card
	TargetPlayer *model.PlayerID
}

func (a *PlayCardTargetOthersAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardTargetOthersAction) IsComplete() bool {
	return a.TargetPlayer != nil
}

func (a *PlayCardTargetOthersAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	return nil
}

func (a *PlayCardTargetOthersAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}

	switch a.Card {
	case model.CardBaron:
		return &inputs.PlayBaronInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	case model.CardPriest:
		return &inputs.PlayPriestInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	case model.CardKing:
		return &inputs.PlayKingInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
		}
	default:
		return nil
	}
}

func (a *PlayCardTargetOthersAction) Type() string {
	return fmt.Sprintf("play_%s", a.Card.Name())
}

func (a *PlayCardTargetOthersAction) String() string {
	if a.TargetPlayer == nil {
		return fmt.Sprintf("play_%s(player=%s)", a.Card.Name(), a.Player)
	}
	return fmt.Sprintf("play_%s(player=%s, target=%s)", a.Card.Name(), a.Player, *a.TargetPlayer)
}

func (a *PlayCardTargetOthersAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayCardTargetOthersAction{
				Player:       a.Player,
				Card:         a.Card,
				TargetPlayer: &player.ID,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}
