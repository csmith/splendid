package model

import (
	"fmt"
)

type PlayerID string

type Game interface {
	PlayerIDs() []PlayerID
}

// Event represents a game event that can be applied to update the game state.
// It is generic over the game type G.
type Event[G Game] interface {
	Typeable
	PlayerID() *PlayerID
	Apply(g G) error
}

// Input represents a game input that can be applied to update the game state.
// It is generic over the game type G.
type Input[G Game] interface {
	Apply(g G, apply func(Event[G])) error
	PlayerID() *PlayerID
}

// Action represents a multi-step player interaction that may require several rounds
// of client-server communication to complete. It is generic over the game type G.
type Action[G Game] interface {
	fmt.Stringer
	Typeable
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(G) []Action[G]
	ToInput() Input[G]
}

// ActionGenerator generates available actions for players based on the current game state.
// It is generic over the game type G.
type ActionGenerator[G Game] interface {
	GenerateActionsForPlayer(g G, playerID PlayerID) []Action[G]
}

// GameUpdate represents the current state of a game along with events and available actions.
// It is generic over the game type G.
type GameUpdate[G Game] struct {
	Game             G
	Event            *Box[Event[G]]
	AvailableActions map[PlayerID]Redactable[[]*Box[Action[G]]]
}
