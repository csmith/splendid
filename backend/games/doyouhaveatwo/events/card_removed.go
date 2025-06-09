package events

import (
	"encoding/json"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventCardRemoved model.EventType = "card_removed"

type CardRemovedEvent struct {
	ResultRemovedCard model.Redactable[model.Card]
}

func (e *CardRemovedEvent) Type() model.EventType {
	return EventCardRemoved
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

func (e *CardRemovedEvent) MarshalJSON() ([]byte, error) {
	type Alias CardRemovedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
