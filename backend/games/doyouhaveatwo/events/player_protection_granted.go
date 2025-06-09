package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerProtectionGranted model.EventType = "player_protection_granted"

type PlayerProtectionGrantedEvent struct {
	Player model.PlayerID `json:"player"`
}

func (e *PlayerProtectionGrantedEvent) Type() model.EventType {
	return EventPlayerProtectionGranted
}

func (e *PlayerProtectionGrantedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerProtectionGrantedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.IsProtected = true
	return nil
}

func (e *PlayerProtectionGrantedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerProtectionGrantedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
