package model

import (
	"github.com/csmith/splendid/backend/model"
)

type PlayerID = model.PlayerID

type Player struct {
	ID          PlayerID                 `json:"id"`
	Name        string                   `json:"name"`
	Hand        []model.Redactable[Card] `json:"hand"`
	DiscardPile []Card                   `json:"discard_pile"`
	TokenCount  int                      `json:"token_count"`
	IsOut       bool                     `json:"is_out"`
	IsProtected bool                     `json:"is_protected"`
}
