package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerEliminated model.EventType = "player_eliminated"

type PlayerEliminatedEvent struct {
	Player               model.PlayerID
	ResultDiscardedCards []model.Card
}

func (e *PlayerEliminatedEvent) Type() model.EventType {
	return EventPlayerEliminated
}

func (e *PlayerEliminatedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerEliminatedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	// Move all cards from hand to discard pile and populate ResultDiscardedCards
	var discardedCards []model.Card
	for _, handCard := range player.Hand {
		discardedCards = append(discardedCards, handCard.Value)
		player.DiscardPile = append(player.DiscardPile, handCard.Value)
	}
	player.Hand = []model.Redactable[model.Card]{}
	e.ResultDiscardedCards = discardedCards

	player.IsOut = true
	return nil
}

func (e *PlayerEliminatedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerEliminatedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
