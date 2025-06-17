package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type CardRemovedEvent struct {
	ResultRemovedCard coremodel.Redactable[model.Card] `json:"removed_card"`
}

func (e *CardRemovedEvent) Type() coremodel.Specifier {
	return specifier("card_removed")
}

func (e *CardRemovedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *CardRemovedEvent) Apply(g *model.Game) error {
	if len(g.Deck) > 0 {
		e.ResultRemovedCard = g.Deck[0]
		g.RemovedCard = &g.Deck[0]
		g.Deck = g.Deck[1:]
	}
	return nil
}
