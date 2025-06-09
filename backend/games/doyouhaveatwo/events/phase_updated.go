package events

import (
	"encoding/json"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventPhaseUpdated model.EventType = "phase_updated"

type PhaseUpdatedEvent struct {
	NewPhase model.GamePhase
}

func (e *PhaseUpdatedEvent) Type() model.EventType {
	return EventPhaseUpdated
}

func (e *PhaseUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *PhaseUpdatedEvent) Apply(g *model.Game) error {
	g.Phase = e.NewPhase
	return nil
}

func (e *PhaseUpdatedEvent) MarshalJSON() ([]byte, error) {
	type Alias PhaseUpdatedEvent
	return json.Marshal(&struct {
		Type model.EventType `json:"type"`
		*Alias
	}{e.Type(), (*Alias)(e)})
}
