package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerTokenAwarded model.EventType = "player_token_awarded"

type PlayerTokenAwardedEvent struct {
	Player model.PlayerID
	Tokens int
}

func (e *PlayerTokenAwardedEvent) Type() model.EventType {
	return EventPlayerTokenAwarded
}

func (e *PlayerTokenAwardedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerTokenAwardedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.TokenCount += e.Tokens
	return nil
}

func (e *PlayerTokenAwardedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerTokenAwardedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
