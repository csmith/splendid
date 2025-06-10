package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPrinceInput struct {
	Player       model.PlayerID
	TargetPlayer model.PlayerID
}

func (i *PlayPrinceInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayPrinceInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Prince card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardPrince,
	})

	// Target player discards their hand
	targetPlayer := g.GetPlayer(i.TargetPlayer)
	if targetPlayer != nil && len(targetPlayer.Hand) > 0 {
		discardedCard := targetPlayer.Hand[0].Value()

		apply(&events.CardDiscardedEvent{
			Player: i.TargetPlayer,
			Card:   discardedCard,
		})

		// If discarded card is Princess, eliminate the target player
		if discardedCard.Value() == model.CardPrincess.Value() {
			apply(&events.PlayerEliminatedEvent{
				Player: i.TargetPlayer,
			})
		} else {
			// Target player draws a new card
			if len(g.Deck) > 0 {
				apply(&events.CardDealtEvent{
					ToPlayer: i.TargetPlayer,
				})
			} else if g.RemovedCard != nil {
				// If deck is empty, give the removed card to the target player
				apply(&events.RemovedCardDealtEvent{
					ToPlayer: i.TargetPlayer,
				})
			}
		}
	}

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
