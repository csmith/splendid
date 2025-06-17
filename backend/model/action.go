package model

import (
	"fmt"
)

// Event represents a game event that can be applied to update the game state.
// It is generic over the game type G.
type Event[G any] interface {
	Typeable
	PlayerID() *PlayerID
	Apply(g G) error
}

// Input represents a game input that can be applied to update the game state.
// It is generic over the game type G.
type Input[G any] interface {
	Apply(g G, apply func(Event[G])) error
	PlayerID() *PlayerID
}

// Action represents a multi-step player interaction that may require several rounds
// of client-server communication to complete. It is generic over the game type G.
type Action[G any] interface {
	fmt.Stringer
	Typeable
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(G) []Action[G]
	ToInput() Input[G]
}

// GameUpdate represents the current state of a game along with events and available actions.
// It is generic over the game type G.
type GameUpdate[G any] struct {
	Game             G
	Event            *Box[Event[G]]
	AvailableActions map[PlayerID]Redactable[[]*Box[Action[G]]]
}
