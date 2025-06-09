package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventPhaseUpdated model.EventType = "phase_updated"

type PhaseUpdatedEvent struct {
	NewPhase model.GamePhase `json:"phase"`
}

func (e *PhaseUpdatedEvent) Type() serialization.Specifier {
	return specifier("phase_updated")
}

func (e *PhaseUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *PhaseUpdatedEvent) Apply(g *model.Game) error {
	g.Phase = e.NewPhase
	return nil
}
