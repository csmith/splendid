package cards

import (
	"errors"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Princess struct{}

func (p Princess) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (p Princess) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (p Princess) Value() int {
	return 8
}

func (p Princess) Name() string {
	return "Princess"
}

func (p Princess) Description() string {
	return "If discarded (played or forced to discard), player is immediately eliminated."
}

func (p Princess) Quantity() int {
	return 1
}