package events

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPlayerDiscardPileCleared model.EventType = "player_discard_pile_cleared"

type PlayerDiscardPileClearedEvent struct {
	Player model.PlayerID
}

func (e *PlayerDiscardPileClearedEvent) Type() model.EventType {
	return EventPlayerDiscardPileCleared
}

func (e *PlayerDiscardPileClearedEvent) PlayerID() *model.PlayerID {
	return &e.Player
}

func (e *PlayerDiscardPileClearedEvent) Apply(g *model.Game) error {
	player := g.GetPlayer(e.Player)
	if player == nil {
		return fmt.Errorf("player with ID %s not found", e.Player)
	}

	player.DiscardPile = []model.Card{}
	return nil
}

func (e *PlayerDiscardPileClearedEvent) MarshalJSON() ([]byte, error) {
	type Alias PlayerDiscardPileClearedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
