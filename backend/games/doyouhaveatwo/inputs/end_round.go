package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputEndRound model.InputType = "end_round"

type EndRoundInput struct {
	Winners []model.PlayerID
}

func (i *EndRoundInput) Type() model.InputType {
	return InputEndRound
}

func (i *EndRoundInput) PlayerID() *model.PlayerID {
	return nil
}

func (i *EndRoundInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Update last round winners
	apply(&events.LastRoundWinnersUpdatedEvent{
		Winners: i.Winners,
	})

	// Award tokens to winners
	for _, winnerID := range i.Winners {
		apply(&events.PlayerTokenAwardedEvent{
			Player: winnerID,
			Tokens: 1,
		})
	}

	// Check if any player has reached the token threshold for game end
	var gameWinners []*model.Player
	for _, player := range g.Players {
		if player.TokenCount >= g.TokensToWin {
			gameWinners = append(gameWinners, player)
		}
	}

	if len(gameWinners) > 0 {
		// Game ends - someone has won
		apply(&events.PhaseUpdatedEvent{
			NewPhase: model.PhaseGameEnd,
		})

		// Declare winner(s) - for simplicity, pick the first one if multiple
		apply(&events.WinnerDeclaredEvent{
			Winner: gameWinners[0].ID,
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
