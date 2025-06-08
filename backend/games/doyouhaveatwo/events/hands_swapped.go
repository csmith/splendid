package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventHandsSwapped model.EventType = "hands_swapped"

type HandsSwappedEvent struct {
	PlayerA     model.PlayerID
	PlayerB     model.PlayerID
	ResultHandA model.Redactable[model.Card]
	ResultHandB model.Redactable[model.Card]
}

func (e *HandsSwappedEvent) Type() model.EventType {
	return EventHandsSwapped
}

func (e *HandsSwappedEvent) PlayerID() *model.PlayerID {
	return &e.PlayerA
}

func (e *HandsSwappedEvent) Apply(g *model.Game) error {
	playerA := g.GetPlayer(e.PlayerA)
	if playerA == nil {
		return fmt.Errorf("player with ID %s not found", e.PlayerA)
	}

	playerB := g.GetPlayer(e.PlayerB)
	if playerB == nil {
		return fmt.Errorf("player with ID %s not found", e.PlayerB)
	}

	if len(playerA.Hand) == 0 || len(playerB.Hand) == 0 {
		return fmt.Errorf("players must have cards to swap hands")
	}

	// Create visibility for both players
	visibility := make(map[model.PlayerID]bool)
	visibility[e.PlayerA] = true
	visibility[e.PlayerB] = true

	// Set result cards visible to both players
	e.ResultHandA = model.Redactable[model.Card]{
		Value:     playerA.Hand[0].Value,
		VisibleTo: visibility,
	}

	e.ResultHandB = model.Redactable[model.Card]{
		Value:     playerB.Hand[0].Value,
		VisibleTo: visibility,
	}

	// Swap the hands
	playerA.Hand, playerB.Hand = playerB.Hand, playerA.Hand

	// Update visibility for swapped cards
	for _, card := range playerA.Hand {
		if card.VisibleTo != nil {
			delete(card.VisibleTo, e.PlayerB)
			card.VisibleTo[e.PlayerA] = true
		}
	}

	for _, card := range playerB.Hand {
		if card.VisibleTo != nil {
			delete(card.VisibleTo, e.PlayerA)
			card.VisibleTo[e.PlayerB] = true
		}
	}

	return nil
}
