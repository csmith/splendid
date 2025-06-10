package inputs

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type DrawCardInput struct {
	Player model.PlayerID
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
		// Deck is empty - trigger showdown
		showdownInput := &ShowdownInput{}
		return showdownInput.Apply(g, apply)
	}

	// Deal card to player
	apply(&events.CardDealtEvent{
		ToPlayer: i.Player,
	})

	// Change phase to play (player can now play a card)
	apply(&events.PhaseUpdatedEvent{
		NewPhase: model.PhasePlay,
	})

	return nil
}
