package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayPrincessInput struct {
	Player model.PlayerID
}

func (i *PlayPrincessInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayPrincessInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Princess card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardPrincess,
	})

	// Playing Princess eliminates the player
	apply(&events.PlayerEliminatedEvent{
		Player: i.Player,
	})

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
