package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type DeckUpdatedEvent struct {
	NewDeck []coremodel.Redactable[model.Card] `json:"deck"`
}

func (e *DeckUpdatedEvent) Type() coremodel.Specifier {
	return specifier("deck_updated")
}

func (e *DeckUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *DeckUpdatedEvent) Apply(g *model.Game) error {
	g.Deck = e.NewDeck
	return nil
}
