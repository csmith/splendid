package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputPlayHandmaid model.InputType = "play_handmaid"

type PlayHandmaidInput struct {
	Player model.PlayerID
}

func (i *PlayHandmaidInput) Type() model.InputType {
	return InputPlayHandmaid
}

func (i *PlayHandmaidInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayHandmaidInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the Handmaid card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardHandmaid,
	})

	// Grant protection to the player
	apply(&events.PlayerProtectionGrantedEvent{
		Player: i.Player,
	})

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
