package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventDrawCard model.EventType = "draw_card"

type DrawCardEvent struct {
	Player    model.PlayerID
	CardDrawn model.Redactable[model.Card]
}

func (e *DrawCardEvent) Type() model.EventType {
	return EventDrawCard
}

func (e *DrawCardEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *DrawCardEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	if len(g.Deck) == 0 {
		return fmt.Errorf("deck is empty, cannot draw card")
	}

	// Draw top card from deck
	card := g.Deck[0]
	g.Deck = g.Deck[1:]

	// Make card visible to the player
	card.VisibleTo[e.Player] = true

	// Add card to player's hand
	player.Hand = append(player.Hand, card)

	e.CardDrawn = card

	return nil
}
