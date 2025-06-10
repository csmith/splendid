package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCountessInput struct {
	Player model.PlayerID
}

func (i *PlayCountessInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayCountessInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Countess card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardCountess,
	})

	// Countess has no special effect

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
