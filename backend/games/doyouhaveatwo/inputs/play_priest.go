package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPriestInput struct {
	Player       model.PlayerID
	TargetPlayer model.PlayerID
}

func (i *PlayPriestInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayPriestInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Priest card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardPriest,
	})

	// Target player reveals their hand to the current player
	apply(&events.HandRevealedEvent{
		SourcePlayer:  i.TargetPlayer,
		TargetPlayers: []model.PlayerID{i.Player},
	})

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
