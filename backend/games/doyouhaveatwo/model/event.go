package model

import (
	coremodel "github.com/csmith/splendid/backend/model"
)

type Event interface {
	coremodel.Typeable
	PlayerID() *PlayerID
	Apply(g *Game) error
}
