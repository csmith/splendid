package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

type HandsSwappedEvent struct {
	PlayerA     model.PlayerID                       `json:"player_a"`
	PlayerB     model.PlayerID                       `json:"player_b"`
	ResultHandA serialization.Redactable[model.Card] `json:"hand_a"`
	ResultHandB serialization.Redactable[model.Card] `json:"hand_b"`
}

func (e *HandsSwappedEvent) Type() coremodel.Specifier {
	return specifier("hands_swapped")
}

func (e *HandsSwappedEvent) PlayerID() *model.PlayerID {
	return &e.PlayerA
}

func (e *HandsSwappedEvent) Apply(g *model.Game) error {
	playerA := g.GetPlayer(e.PlayerA)
	if playerA == nil {
		return fmt.Errorf("player with ID %s not found", e.PlayerA)
	}

	playerB := g.GetPlayer(e.PlayerB)
	if playerB == nil {
		return fmt.Errorf("player with ID %s not found", e.PlayerB)
	}

	if len(playerA.Hand) == 0 || len(playerB.Hand) == 0 {
		return fmt.Errorf("players must have cards to swap hands")
	}

	// Set result cards visible to both players
	e.ResultHandA = serialization.NewRedactable(playerA.Hand[0].Value(), e.PlayerA, e.PlayerB)
	e.ResultHandB = serialization.NewRedactable(playerB.Hand[0].Value(), e.PlayerA, e.PlayerB)

	// Swap the hands
	playerA.Hand, playerB.Hand = playerB.Hand, playerA.Hand

	// Update visibility for swapped cards
	for i, card := range playerA.Hand {
		playerA.Hand[i] = card.WithVisibility(e.PlayerA)
	}

	for i, card := range playerB.Hand {
		playerB.Hand[i] = card.WithVisibility(e.PlayerB)
	}

	return nil
}
