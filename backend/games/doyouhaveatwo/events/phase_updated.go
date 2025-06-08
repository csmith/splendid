package events

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

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
