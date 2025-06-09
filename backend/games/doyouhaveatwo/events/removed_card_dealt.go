package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventRemovedCardDealt model.EventType = "removed_card_dealt"

type RemovedCardDealtEvent struct {
	ToPlayer               model.PlayerID                       `json:"to_player"`
	ResultRemovedCardDealt serialization.Redactable[model.Card] `json:"removed_card_dealt"`
}

func (e *RemovedCardDealtEvent) Type() serialization.Specifier {
	return specifier("removed_card_dealt")
}

func (e *RemovedCardDealtEvent) PlayerID() *model.PlayerID {
	return &e.ToPlayer
}

func (e *RemovedCardDealtEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.ToPlayer)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.ToPlayer)
	}

	if g.RemovedCard == nil {
		return fmt.Errorf("no removed card available")
	}

	// Take the removed card
	card := *g.RemovedCard

	// Make card visible to the player
	card = card.WithVisibility(e.ToPlayer)

	// Add to player's hand
	player.Hand = append(player.Hand, card)

	// Set result for client visibility
	e.ResultRemovedCardDealt = card

	// Clear the removed card
	g.RemovedCard = nil

	return nil
}
