package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerProtectionCleared model.EventType = "player_protection_cleared"

type PlayerProtectionClearedEvent struct {
	Player model.PlayerID
}

func (e *PlayerProtectionClearedEvent) Type() model.EventType {
	return EventPlayerProtectionCleared
}

func (e *PlayerProtectionClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerProtectionClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsProtected = false
	return nil
}

func (e *PlayerProtectionClearedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerProtectionClearedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
