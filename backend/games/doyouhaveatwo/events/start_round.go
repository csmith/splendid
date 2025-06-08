package events

import (
	"math/rand"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventStartRound model.EventType = "start_round"

type StartRoundEvent struct {
	DeckShuffled *[]model.Redactable[model.Card]                  `json:",omitempty"`
	RemovedCard  *model.Redactable[model.Card]                    `json:",omitempty"`
	DealtCards   map[model.PlayerID]*model.Redactable[model.Card] `json:",omitempty"`
}

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
	deck := e.createShuffledDeck()
	e.DeckShuffled = &deck
	g.Deck = deck

	// Remove top card
	removedCard := g.Deck[0]
	e.RemovedCard = &removedCard
	g.RemovedCard = &removedCard
	g.Deck = g.Deck[1:]

	// Deal one card to each player
	e.DealtCards = make(map[model.PlayerID]*model.Redactable[model.Card])
	for _, player := range g.Players {
		if len(g.Deck) > 0 {
			card := g.Deck[0]
			g.Deck = g.Deck[1:]
			card.VisibleTo[player.ID] = true
			player.Hand = []model.Redactable[model.Card]{card}
			e.DealtCards[player.ID] = &card
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
	for _, cardType := range model.CardTypes {
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
