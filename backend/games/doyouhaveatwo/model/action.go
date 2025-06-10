package model

import (
	"fmt"

	"github.com/csmith/splendid/backend/serialization"
)

type Input interface {
	Apply(g *Game, apply func(Event)) error
	PlayerID() *PlayerID
}

type Action interface {
	fmt.Stringer
	serialization.Typeable
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(*Game) []Action
	ToInput() Input
}

type GameUpdate struct {
	Game             Game
	Event            *serialization.Box[Event]
	AvailableActions map[PlayerID]serialization.Redactable[[]*serialization.Box[Action]]
}
