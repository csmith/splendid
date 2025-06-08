package cards

import (
	"errors"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Prince struct{}

func (p Prince) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (p Prince) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (p Prince) Value() int {
	return 5
}

func (p Prince) Name() string {
	return "Prince"
}

func (p Prince) Description() string {
	return "Target player discards their hand and draws a new card. Can target self."
}

func (p Prince) Quantity() int {
	return 2
}