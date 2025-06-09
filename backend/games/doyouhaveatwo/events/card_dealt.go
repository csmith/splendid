package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventCardDealt model.EventType = "card_dealt"

type CardDealtEvent struct {
	ToPlayer        model.PlayerID               `json:"to_player"`
	ResultCardDealt model.Redactable[model.Card] `json:"card_dealt"`
}

func (e *CardDealtEvent) Type() serialization.Specifier {
	return specifier("card_dealt")
}

func (e *CardDealtEvent) PlayerID() *model.PlayerID {
	return &e.ToPlayer
}

func (e *CardDealtEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.ToPlayer)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.ToPlayer)
	}

	if len(g.Deck) == 0 {
		return fmt.Errorf("deck is empty, cannot deal card")
	}

	// Take top card from deck
	card := g.Deck[0]
	g.Deck = g.Deck[1:]

	// Make card visible to the player
	card = card.WithVisibility(e.ToPlayer)

	// Add to player's hand
	player.Hand = append(player.Hand, card)

	// Set result for client visibility
	e.ResultCardDealt = card

	return nil
}
