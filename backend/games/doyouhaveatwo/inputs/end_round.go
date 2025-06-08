package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputEndRound InputType = "end_round"

type EndRoundInput struct {
	Player model.PlayerID
}

func (i *EndRoundInput) Type() InputType {
	return InputEndRound
}

func (i *EndRoundInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *EndRoundInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Find players who have reached the token threshold
	var winners []*model.Player
	for _, player := range g.Players {
		if player.TokenCount >= g.TokensToWin {
			winners = append(winners, player)
		}
	}

	if len(winners) > 0 {
		// Game ends - someone has won
		apply(&events.PhaseUpdatedEvent{
			NewPhase: model.PhaseGameEnd,
		})

		// Declare winner(s) - for simplicity, pick the first one if multiple
		// In a real implementation, you might want to handle ties differently
		apply(&events.WinnerDeclaredEvent{
			Winner: winners[0].ID,
		})
	} else {
		// No winner yet - continue to next round
		apply(&events.PhaseUpdatedEvent{
			NewPhase: model.PhaseRoundEnd,
		})

		// Start the next round
		startRoundInput := &StartRoundInput{}
		return startRoundInput.Apply(g, apply)
	}

	return nil
}
