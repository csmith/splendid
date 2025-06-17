package model

import (
	coremodel "github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

type Input interface {
	Apply(g *Game, apply func(Event)) error
	PlayerID() *PlayerID
}

// Action is a type alias to the generic Action interface from the shared model package
type Action = coremodel.Action[*Game, Input]

// GameAction is a convenience type alias for actions in this specific game
type GameAction = Action

type GameUpdate struct {
	Game             Game
	Event            *serialization.Box[Event]
	AvailableActions map[PlayerID]serialization.Redactable[[]*serialization.Box[GameAction]]
}
