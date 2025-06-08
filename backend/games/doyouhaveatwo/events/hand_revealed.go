package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const EventHandRevealed model.EventType = "hand_revealed"

type HandRevealedEvent struct {
	SourcePlayer       model.PlayerID
	TargetPlayers      []model.PlayerID
	ResultRevealedCard model.Redactable[model.Card]
}

func (e *HandRevealedEvent) Type() model.EventType {
	return EventHandRevealed
}

func (e *HandRevealedEvent) PlayerID() *model.PlayerID {
	return &e.SourcePlayer
}

func (e *HandRevealedEvent) Apply(g *model.Game) error {
	sourcePlayer := g.GetPlayer(e.SourcePlayer)
	if sourcePlayer == nil {
		return fmt.Errorf("source player with ID %s not found", e.SourcePlayer)
	}

	if len(sourcePlayer.Hand) == 0 {
		return fmt.Errorf("source player %s has no cards to reveal", e.SourcePlayer)
	}

	// Create a new redacted card visible to source and all target players
	card := sourcePlayer.Hand[0].Value
	visibility := make(map[model.PlayerID]bool)
	visibility[e.SourcePlayer] = true
	for _, targetPlayer := range e.TargetPlayers {
		visibility[targetPlayer] = true
	}

	// Set result for client visibility
	e.ResultRevealedCard = model.Redactable[model.Card]{
		Value:     card,
		VisibleTo: visibility,
	}

	return nil
}
