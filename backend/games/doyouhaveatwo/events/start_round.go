package events

import (
	"math/rand"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/cards"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventStartRound model.EventType = "start_round"

type StartRoundEvent struct{}

func (e *StartRoundEvent) Type() model.EventType {
	return EventStartRound
}

func (e *StartRoundEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *StartRoundEvent) Apply(g *model.Game) error {
	// Increment round number
	g.Round++

	// Create new shuffled deck
	g.Deck = e.createShuffledDeck()

	// Remove top card
	g.RemovedCard = g.Deck[0]
	g.Deck = g.Deck[1:]

	// Deal one card to each player
	for _, player := range g.Players {
		if len(g.Deck) > 0 {
			card := g.Deck[0]
			g.Deck = g.Deck[1:]
			card.VisibleTo[player.ID] = true
			player.Hand = []model.Redactable[model.Card]{card}
		}
	}

	// Reset player states
	for _, player := range g.Players {
		player.IsOut = false
		player.IsProtected = false
		player.DiscardPile = []model.Card{}
	}

	// Set phase to play
	g.Phase = model.PhasePlay

	return nil
}

func (e *StartRoundEvent) createShuffledDeck() []model.Redactable[model.Card] {
	var deck []model.Redactable[model.Card]

	// Add cards according to their quantities
	cardTypes := []model.Card{
		cards.Guard{}, cards.Priest{}, cards.Baron{}, cards.Handmaid{},
		cards.Prince{}, cards.King{}, cards.Countess{}, cards.Princess{},
	}

	for _, cardType := range cardTypes {
		for i := 0; i < cardType.Quantity(); i++ {
			deck = append(deck, model.Redactable[model.Card]{
				Value:     cardType,
				VisibleTo: make(map[model.PlayerID]bool),
			})
		}
	}

	// Shuffle the deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}