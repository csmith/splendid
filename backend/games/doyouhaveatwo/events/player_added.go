package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventPlayerAdded model.EventType = "player_added"

type PlayerAddedEvent struct {
	NewPlayerID   model.PlayerID `json:"player"`
	NewPlayerName string         `json:"name"`
	Position      int            `json:"position"`
}

func (e *PlayerAddedEvent) Type() serialization.Specifier {
	return specifier("player_added")
}

func (e *PlayerAddedEvent) PlayerID() *model.PlayerID {
	return &e.NewPlayerID
}

func (e *PlayerAddedEvent) Apply(g *model.Game) error {
	// Create new player
	newPlayer := &model.Player{
		ID:            e.NewPlayerID,
		Name:          e.NewPlayerName,
		Hand:          []serialization.Redactable[model.Card]{},
		DiscardPile:   []model.Card{},
		TokenCount:    0,
		IsOut:         false,
		IsProtected:   false,
		Position:      e.Position,
		PendingAction: serialization.Redactable[*serialization.Box[model.Action]]{},
	}

	// Add player to game
	g.Players = append(g.Players, newPlayer)

	return nil
}
