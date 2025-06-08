package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputPlayKing model.InputType = "play_king"

type PlayKingInput struct {
	Player       model.PlayerID
	TargetPlayer model.PlayerID
}

func (i *PlayKingInput) Type() model.InputType {
	return InputPlayKing
}

func (i *PlayKingInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *PlayKingInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Discard the King card from player's hand
	apply(&events.CardDiscardedEvent{
		Player: i.Player,
		Card:   model.CardKing,
	})

	// Trade hands between current player and target player
	apply(&events.HandsSwappedEvent{
		PlayerA: i.Player,
		PlayerB: i.TargetPlayer,
	})

	// End the turn
	endTurnInput := &EndTurnInput{
		Player: i.Player,
	}

	return endTurnInput.Apply(g, apply)
}
