package inputs

import (
	"math/rand"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type StartRoundInput struct{}

func (i *StartRoundInput) PlayerID() *model.PlayerID {
	return nil
}

func (i *StartRoundInput) Apply(g *model.Game, apply func(model.Event)) error {
	for _, player := range g.Players {
		if player.IsOut {
			apply(&events.PlayerRestoredEvent{
				Player: player.ID,
			})
		}
		if player.IsProtected {
			apply(&events.PlayerProtectionClearedEvent{
				Player: player.ID,
			})
		}
		if len(player.DiscardPile) > 0 {
			apply(&events.PlayerDiscardPileClearedEvent{
				Player: player.ID,
			})
		}
		if len(player.Hand) > 0 {
			apply(&events.PlayerHandClearedEvent{
				Player: player.ID,
			})
		}
	}

	// Increment round
	apply(&events.RoundUpdatedEvent{
		NewRound: g.Round + 1,
	})

	// Set current player: random winner from last round, or random if no winners
	var currentPlayerIndex int
	if len(g.LastRoundWinners) > 0 {
		randomWinner := g.LastRoundWinners[rand.Intn(len(g.LastRoundWinners))]
		// Find the index of the winner
		for i, player := range g.Players {
			if player.ID == randomWinner {
				currentPlayerIndex = i
				break
			}
		}
	} else {
		// No winners recorded, pick random player
		currentPlayerIndex = rand.Intn(len(g.Players))
	}

	apply(&events.CurrentPlayerUpdatedEvent{
		NewCurrentPlayer: currentPlayerIndex,
	})

	// Create and set new shuffled deck
	deck := i.createShuffledDeck()
	apply(&events.DeckUpdatedEvent{
		NewDeck: deck,
	})

	// Remove top card from deck
	apply(&events.CardRemovedEvent{})

	// For two-player games, set aside 3 cards that are visible to everyone
	if len(g.Players) == 2 {
		apply(&events.CardsSetAsideEvent{})
	}

	// Deal cards to all players
	for _, player := range g.Players {
		apply(&events.CardDealtEvent{
			ToPlayer: player.ID,
		})
	}

	// Set phase to draw (first player needs to draw a card)
	apply(&events.PhaseUpdatedEvent{
		NewPhase: model.PhaseDraw,
	})

	return nil
}

func (i *StartRoundInput) createShuffledDeck() []serialization.Redactable[model.Card] {
	var deck []serialization.Redactable[model.Card]

	// Add cards according to their quantities
	for _, cardType := range model.CardTypes {
		for j := 0; j < cardType.Quantity(); j++ {
			deck = append(deck, serialization.NewRedactable(cardType))
		}
	}

	// Shuffle the deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
