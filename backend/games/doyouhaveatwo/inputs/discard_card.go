package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputDiscardCard model.InputType = "discard_card"

type DiscardCardInput struct {
	Player   model.PlayerID
	CardName string
}

func (i *DiscardCardInput) Type() model.InputType {
	return InputDiscardCard
}

func (i *DiscardCardInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *DiscardCardInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Find the card by name
	card, err := model.GetCardByName(i.CardName)
	if err != nil {
		return err
	}

	// Discard the specified card from player's hand (no effect)
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   card,
	})

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
