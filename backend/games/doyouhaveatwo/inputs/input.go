package inputs

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

type InputType string

type Input interface {
	Apply(g *model.Game) ([]model.Event, error)
	Type() InputType
	PlayerID() *model.PlayerID
}
