package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventCardDiscarded model.EventType = "card_discarded"

type CardDiscardedEvent struct {
	Player              model.PlayerID `json:"player"`
	Card                model.Card     `json:"card"`
	ResultCardDiscarded model.Card     `json:"discarded_card"`
}

func (e *CardDiscardedEvent) Type() serialization.Specifier {
	return specifier("card_discarded")
}

func (e *CardDiscardedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *CardDiscardedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	// Find and remove the card from player's hand
	cardIndex := -1
	for i, handCard := range player.Hand {
		if handCard.Value().Value() == e.Card.Value() && handCard.Value().Name() == e.Card.Name() {
			cardIndex = i
			break
		}
	}

	if cardIndex == -1 {
		return fmt.Errorf("card %s not found in player %s hand", e.Card.Name(), e.Player)
	}

	// Remove card from hand
	player.Hand = append(player.Hand[:cardIndex], player.Hand[cardIndex+1:]...)

	// Add card to discard pile (discarded cards are public)
	player.DiscardPile = append(player.DiscardPile, e.Card)

	// Set result for client visibility
	e.ResultCardDiscarded = e.Card

	return nil
}
