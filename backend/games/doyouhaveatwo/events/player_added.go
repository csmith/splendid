package events

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerAdded model.EventType = "player_added"

type PlayerAddedEvent struct {
	NewPlayerID   model.PlayerID `json:"player"`
	NewPlayerName string         `json:"name"`
	Position      int            `json:"position"`
}

func (e *PlayerAddedEvent) Type() model.EventType {
	return EventPlayerAdded
}

func (e *PlayerAddedEvent) PlayerID() *model.PlayerID {
	return &e.NewPlayerID
}

func (e *PlayerAddedEvent) Apply(g *model.Game) error {
	// Create new player
	newPlayer := &model.Player{
		ID:            e.NewPlayerID,
		Name:          e.NewPlayerName,
		Hand:          []model.Redactable[model.Card]{},
		DiscardPile:   []model.Card{},
		TokenCount:    0,
		IsOut:         false,
		IsProtected:   false,
		Position:      e.Position,
		PendingAction: model.Redactable[model.Action]{},
	}

	// Add player to game
	g.Players = append(g.Players, newPlayer)

	return nil
}

func (e *PlayerAddedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerAddedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
