package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputPlayBaron model.InputType = "play_baron"

type PlayBaronInput struct {
	Player       model.PlayerID
	TargetPlayer model.PlayerID
}

func (i *PlayBaronInput) Type() model.InputType {
	return InputPlayBaron
}

func (i *PlayBaronInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayBaronInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Baron card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardBaron,
	})

	// Both players reveal their hands to each other
	apply(&events.HandRevealedEvent{
		SourcePlayer:  i.Player,
		TargetPlayers: []model.PlayerID{i.TargetPlayer},
	})

	apply(&events.HandRevealedEvent{
		SourcePlayer:  i.TargetPlayer,
		TargetPlayers: []model.PlayerID{i.Player},
	})

	// Compare hands and eliminate lower value player
	currentPlayer := g.GetPlayer(i.Player)
	targetPlayer := g.GetPlayer(i.TargetPlayer)

	if currentPlayer != nil && targetPlayer != nil &&
		len(currentPlayer.Hand) > 0 && len(targetPlayer.Hand) > 0 {

		currentCard := currentPlayer.Hand[0].Value()
		targetCard := targetPlayer.Hand[0].Value()

		if currentCard.Value() < targetCard.Value() {
			// Current player has lower value, eliminate them
			apply(&events.PlayerEliminatedEvent{
				Player: i.Player,
			})
		} else if targetCard.Value() < currentCard.Value() {
			// Target player has lower value, eliminate them
			apply(&events.PlayerEliminatedEvent{
				Player: i.TargetPlayer,
			})
		}
		// If tied, nothing happens
	}

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
