package model

type PlayerID string

type Player struct {
	ID            PlayerID           `json:"id"`
	Name          string             `json:"name"`
	Hand          []Redactable[Card] `json:"hand"`
	DiscardPile   []Card             `json:"discard_pile"`
	TokenCount    int                `json:"token_count"`
	IsOut         bool               `json:"is_out"`
	IsProtected   bool               `json:"is_protected"`
	Position      int                `json:"position"`
	PendingAction Redactable[Action] `json:"pending_action"`
}
