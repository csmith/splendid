package inputs

import (
	"math/rand"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputStartRound model.InputType = "start_round"

type StartRoundInput struct{}

func (i *StartRoundInput) Type() model.InputType {
	return InputStartRound
}

func (i *StartRoundInput) PlayerID() *model.PlayerID {
	return nil
}

func (i *StartRoundInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Increment round
	apply(&events.RoundUpdatedEvent{
		NewRound: g.Round + 1,
	})

	// Create and set new shuffled deck
	deck := i.createShuffledDeck()
	apply(&events.DeckUpdatedEvent{
		NewDeck: deck,
	})

	// Remove top card from deck
	apply(&events.CardRemovedEvent{})

	// Deal cards to all players
	for _, player := range g.Players {
		apply(&events.CardDealtEvent{
			ToPlayer: player.ID,
		})
	}

	// Reset all player states
	for _, player := range g.Players {
		apply(&events.PlayerRestoredEvent{
			Player: player.ID,
		})
		apply(&events.PlayerProtectionClearedEvent{
			Player: player.ID,
		})
		apply(&events.PlayerDiscardPileClearedEvent{
			Player: player.ID,
		})
	}

	// Set phase to play
	apply(&events.PhaseUpdatedEvent{
		NewPhase: model.PhasePlay,
	})

	return nil
}

func (i *StartRoundInput) createShuffledDeck() []model.Redactable[model.Card] {
	var deck []model.Redactable[model.Card]

	// Add cards according to their quantities
	for _, cardType := range model.CardTypes {
		for j := 0; j < cardType.Quantity(); j++ {
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
