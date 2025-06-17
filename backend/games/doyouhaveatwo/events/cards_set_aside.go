package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type CardsSetAsideEvent struct {
	ResultSetAsideCards []model.Card `json:"set_aside_cards"`
}

func (e *CardsSetAsideEvent) Type() coremodel.Specifier {
	return specifier("cards_set_aside")
}

func (e *CardsSetAsideEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *CardsSetAsideEvent) Apply(g *model.Game) error {
	if len(g.Deck) < 3 {
		return fmt.Errorf("cannot set aside 3 cards: deck only has %d cards", len(g.Deck))
	}

	var cards []model.Card
	for i := 0; i < 3; i++ {
		cards = append(cards, g.Deck[i].Value())
	}

	e.ResultSetAsideCards = cards
	g.SetAsideCards = cards
	g.Deck = g.Deck[3:]
	return nil
}
