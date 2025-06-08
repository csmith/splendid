package inputs

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputDrawCard model.InputType = "draw_card"

type DrawCardInput struct {
	Player model.PlayerID
}

func (i *DrawCardInput) Type() model.InputType {
	return InputDrawCard
}

func (i *DrawCardInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *DrawCardInput) Apply(g *model.Game, apply func(model.Event)) error {
	player := g.GetPlayer(i.Player)
	if player == nil {
		return fmt.Errorf("player not found: %s", i.Player)
	}

	if len(g.Deck) == 0 {
		return fmt.Errorf("deck is empty")
	}

	// Deal card to player
	apply(&events.CardDealtEvent{
		ToPlayer: i.Player,
	})

	return nil
}
