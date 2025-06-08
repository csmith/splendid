package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputDrawCard InputType = "draw_card"

type DrawCardInput struct {
	Player model.PlayerID
}

func (i *DrawCardInput) Type() InputType {
	return InputDrawCard
}

func (i *DrawCardInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *DrawCardInput) Apply(g *model.Game) ([]model.Event, error) {
	var eventList []model.Event

	// Validate player exists
	player := g.GetPlayer(i.Player)
	if player == nil {
		// Return empty event list for invalid input (no-op)
		return eventList, nil
	}

	// Validate deck has cards
	if len(g.Deck) == 0 {
		// Return empty event list for invalid input (no-op)
		return eventList, nil
	}

	// Deal card to player
	eventList = append(eventList, &events.CardDealtEvent{
		ToPlayer: i.Player,
	})

	return eventList, nil
}
