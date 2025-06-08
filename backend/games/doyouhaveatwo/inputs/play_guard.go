package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputPlayGuard InputType = "play_guard"

type PlayGuardInput struct {
	Player       model.PlayerID
	TargetPlayer model.PlayerID
	GuessedRank  int
}

func (i *PlayGuardInput) Type() InputType {
	return InputPlayGuard
}

func (i *PlayGuardInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayGuardInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Guard card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardGuard,
	})

	// Check if guess is correct
	targetPlayer := g.GetPlayer(i.TargetPlayer)
	if targetPlayer != nil && len(targetPlayer.Hand) > 0 {
		// Check if any card in target's hand matches the guessed rank
		for _, handCard := range targetPlayer.Hand {
			if handCard.Value.Value() == i.GuessedRank {
				// Guess is correct - eliminate the target player (applied immediately!)
				apply(&events.PlayerEliminatedEvent{
					Player: i.TargetPlayer,
				})
				break
			}
		}
	}

	// End the turn (now sees the correct state with player eliminated)
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
