package events

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventLastRoundWinnersUpdated model.EventType = "last_round_winners_updated"

type LastRoundWinnersUpdatedEvent struct {
	Winners []model.PlayerID `json:"winners"`
}

func (e *LastRoundWinnersUpdatedEvent) Type() model.EventType {
	return EventLastRoundWinnersUpdated
}

func (e *LastRoundWinnersUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *LastRoundWinnersUpdatedEvent) Apply(g *model.Game) error {
	g.LastRoundWinners = e.Winners
	return nil
}

func (e *LastRoundWinnersUpdatedEvent) MarshalJSON() ([]byte, error) {
	type Alias LastRoundWinnersUpdatedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
