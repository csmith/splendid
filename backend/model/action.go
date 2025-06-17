package model

import (
	"fmt"
)

// Input represents a game input that can be applied to update the game state.
// It is generic over the game type G and event type E.
type Input[G any, E any] interface {
	Apply(g G, apply func(E)) error
	PlayerID() *PlayerID
}

// Action represents a multi-step player interaction that may require several rounds
// of client-server communication to complete. It is generic over the game type G
// and event type E, allowing it to be used by different game implementations.
type Action[G any, E any] interface {
	fmt.Stringer
	Typeable
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(G) []Action[G, E]
	ToInput() Input[G, E]
}
