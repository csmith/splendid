package model

type PlayerID string

type Player struct {
	ID          PlayerID
	Name        string
	Hand        []Redactable[Card]
	DiscardPile []Card
	TokenCount  int
	IsOut       bool
	IsProtected bool
	Position    int
}
