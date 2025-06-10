package model

import (
	"github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayerID = model.PlayerID

type Player struct {
	ID            PlayerID                                             `json:"id"`
	Name          string                                               `json:"name"`
	Hand          []serialization.Redactable[Card]                     `json:"hand"`
	DiscardPile   []Card                                               `json:"discard_pile"`
	TokenCount    int                                                  `json:"token_count"`
	IsOut         bool                                                 `json:"is_out"`
	IsProtected   bool                                                 `json:"is_protected"`
	PendingAction serialization.Redactable[*serialization.Box[Action]] `json:"pending_action"`
}
