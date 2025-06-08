package events

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

const EventDeckUpdated model.EventType = "deck_updated"

type DeckUpdatedEvent struct {
	NewDeck []model.Redactable[model.Card]
}

func (e *DeckUpdatedEvent) Type() model.EventType {
	return EventDeckUpdated
}

func (e *DeckUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *DeckUpdatedEvent) Apply(g *model.Game) error {
	g.Deck = e.NewDeck
	return nil
}
