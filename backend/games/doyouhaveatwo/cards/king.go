package cards

import (
	"errors"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type King struct{}

func (k King) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (k King) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (k King) Value() int {
	return 6
}

func (k King) Name() string {
	return "King"
}

func (k King) Description() string {
	return "Trade hands with another player."
}

func (k King) Quantity() int {
	return 1
}
