package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventDeckUpdated model.EventType = "deck_updated"

type DeckUpdatedEvent struct {
	NewDeck []model.Redactable[model.Card] `json:"deck"`
}

func (e *DeckUpdatedEvent) Type() serialization.Specifier {
	return specifier("deck_updated")
}

func (e *DeckUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *DeckUpdatedEvent) Apply(g *model.Game) error {
	g.Deck = e.NewDeck
	return nil
}
