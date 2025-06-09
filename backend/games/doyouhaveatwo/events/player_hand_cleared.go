package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerHandCleared model.EventType = "player_hand_cleared"

type PlayerHandClearedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerHandClearedEvent) Type() model.EventType {
	return EventPlayerHandCleared
}

func (e *PlayerHandClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerHandClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.Hand = []model.Redactable[model.Card]{}
	return nil
}

func (e *PlayerHandClearedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerHandClearedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
