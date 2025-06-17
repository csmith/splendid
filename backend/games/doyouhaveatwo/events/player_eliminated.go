package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerEliminatedEvent struct {
	Player               model.PlayerID `json:"player"`
	ResultDiscardedCards []model.Card   `json:"discarded_cards"`
}

func (e *PlayerEliminatedEvent) Type() coremodel.Specifier {
	return specifier("player_eliminated")
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
		discardedCards = append(discardedCards, handCard.Value())
		player.DiscardPile = append(player.DiscardPile, handCard.Value())
	}
	player.Hand = []serialization.Redactable[model.Card]{}
	e.ResultDiscardedCards = discardedCards

	player.IsOut = true
	return nil
}
