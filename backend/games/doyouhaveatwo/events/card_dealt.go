package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventCardDealt model.EventType = "card_dealt"

type CardDealtEvent struct {
	ToPlayer        model.PlayerID
	ResultCardDealt model.Redactable[model.Card]
}

func (e *CardDealtEvent) Type() model.EventType {
	return EventCardDealt
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
	card.VisibleTo[e.ToPlayer] = true

	// Add to player's hand
	player.Hand = append(player.Hand, card)

	// Set result for client visibility
	e.ResultCardDealt = card

	return nil
}

func (e *CardDealtEvent) MarshalJSON() ([]byte, error) {
	type Alias CardDealtEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
