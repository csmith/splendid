package model

import (
	coremodel "github.com/csmith/splendid/backend/model"
	"github.com/csmith/splendid/backend/serialization"
)

// Event is a type alias to the generic Event interface from the shared model package
type Event = coremodel.Event[*Game]

// Input is a type alias to the generic Input interface from the shared model package
type Input = coremodel.Input[*Game]

// Action is a type alias to the generic Action interface from the shared model package
type Action = coremodel.Action[*Game]

// GameAction is a convenience type alias for actions in this specific game
type GameAction = Action

type GameUpdate struct {
	Game             Game
	Event            *serialization.Box[Event]
	AvailableActions map[PlayerID]serialization.Redactable[[]*serialization.Box[GameAction]]
}
