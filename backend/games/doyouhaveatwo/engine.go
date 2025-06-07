package doyouhaveatwo

import (
	"math/rand"
)

type Engine struct {
	Game       Game
	updateChan chan<- GameUpdate
	eventChan  <-chan Event
}

func (e *Engine) applyEvent(event Event) {
	switch event.Type {
	case EventStartRound:
		e.startRound()
	}

	e.updateChan <- GameUpdate{
		Game:             e.Game,
		Event:            event,
		AvailableActions: make(map[PlayerID]Redactable[[]Action]),
	}
}

func (e *Engine) startRound() {
	// Increment round number
	e.Game.Round++

	// Create new deck
	e.Game.Deck = e.createDeck()

	// Shuffle deck
	e.shuffleDeck()

	// Remove top card
	e.Game.RemovedCard = e.Game.Deck[0]
	e.Game.Deck = e.Game.Deck[1:]

	// Deal one card to each player
	for _, player := range e.Game.Players {
		if len(e.Game.Deck) > 0 {
			card := e.Game.Deck[0]
			e.Game.Deck = e.Game.Deck[1:]
			card.VisibleTo[player.ID] = true
			player.Hand = []Redactable[Card]{card}
		}
	}

	// Reset player states
	for _, player := range e.Game.Players {
		player.IsOut = false
		player.IsProtected = false
		player.DiscardPile = []Card{}
	}

	// Set phase to play
	e.Game.Phase = PhasePlay
}

func (e *Engine) createDeck() []Redactable[Card] {
	var deck []Redactable[Card]

	// Add cards according to their quantities
	cardTypes := []Card{
		Guard{}, Priest{}, Baron{}, Handmaid{},
		Prince{}, King{}, Countess{}, Princess{},
	}

	for _, cardType := range cardTypes {
		for i := 0; i < cardType.Quantity(); i++ {
			deck = append(deck, Redactable[Card]{
				Value:     cardType,
				VisibleTo: make(map[PlayerID]bool),
			})
		}
	}

	return deck
}

func (e *Engine) shuffleDeck() {
	rand.Shuffle(len(e.Game.Deck), func(i, j int) {
		e.Game.Deck[i], e.Game.Deck[j] = e.Game.Deck[j], e.Game.Deck[i]
	})
}
