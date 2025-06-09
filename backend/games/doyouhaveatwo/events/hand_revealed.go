package events

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

const EventHandRevealed model.EventType = "hand_revealed"

type HandRevealedEvent struct {
	SourcePlayer       model.PlayerID                       `json:"source_player"`
	TargetPlayers      []model.PlayerID                     `json:"target_players"`
	ResultRevealedCard serialization.Redactable[model.Card] `json:"revealed_card"`
}

func (e *HandRevealedEvent) Type() serialization.Specifier {
	return specifier("hand_revealed")
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

	players := append([]model.PlayerID{e.SourcePlayer}, e.TargetPlayers...)

	e.ResultRevealedCard = sourcePlayer.Hand[0].WithVisibility(players...)

	return nil
}
