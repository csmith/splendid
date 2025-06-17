package model

import (
	"github.com/csmith/splendid/backend/serialization"
)

type Event interface {
	serialization.Typeable
	PlayerID() *PlayerID
	Apply(g *Game) error
}
