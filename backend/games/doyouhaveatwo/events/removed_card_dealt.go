package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventRemovedCardDealt model.EventType = "removed_card_dealt"

type RemovedCardDealtEvent struct {
	ToPlayer               model.PlayerID
	ResultRemovedCardDealt model.Redactable[model.Card]
}

func (e *RemovedCardDealtEvent) Type() model.EventType {
	return EventRemovedCardDealt
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
	card.VisibleTo[e.ToPlayer] = true

	// Add to player's hand
	player.Hand = append(player.Hand, card)

	// Set result for client visibility
	e.ResultRemovedCardDealt = card

	// Clear the removed card
	g.RemovedCard = nil

	return nil
}
