package doyouhaveatwo

import (
	"fmt"
)

type EventType string

const (
	EventStartRound EventType = "start_round"
	EventDrawCard   EventType = "draw_card"
)

type Event struct {
	Type     EventType
	PlayerID *PlayerID
	Details  interface{}
}

func (t EventType) Apply(e *Event, g *Game) error {
	switch t {
	case EventStartRound:
		return applyStartRound(g)
	case EventDrawCard:
		return applyDrawCard(e, g)
	default:
		return fmt.Errorf("unknown event type: %s", t)
	}
}

func applyStartRound(g *Game) error {
	// Increment round number
	g.Round++

	// Create new shuffled deck
	g.Deck = CreateShuffledDeck()

	// Remove top card
	g.RemovedCard = g.Deck[0]
	g.Deck = g.Deck[1:]

	// Deal one card to each player
	for _, player := range g.Players {
		if len(g.Deck) > 0 {
			card := g.Deck[0]
			g.Deck = g.Deck[1:]
			card.VisibleTo[player.ID] = true
			player.Hand = []Redactable[Card]{card}
		}
	}

	// Reset player states
	for _, player := range g.Players {
		player.IsOut = false
		player.IsProtected = false
		player.DiscardPile = []Card{}
	}

	// Set phase to play
	g.Phase = PhasePlay

	return nil
}

func applyDrawCard(e *Event, g *Game) error {
	if e.PlayerID == nil {
		return fmt.Errorf("player ID cannot be nil")
	}

	player := g.GetPlayer(*e.PlayerID)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", *e.PlayerID)
	}

	if len(g.Deck) == 0 {
		return fmt.Errorf("deck is empty, cannot draw card")
	}

	// Draw top card from deck
	card := g.Deck[0]
	g.Deck = g.Deck[1:]

	// Make card visible to the player
	card.VisibleTo[*e.PlayerID] = true

	// Add card to player's hand
	player.Hand = append(player.Hand, card)

	return nil
}

