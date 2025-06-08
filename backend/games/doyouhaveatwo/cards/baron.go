package cards

import (
	"errors"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Baron struct{}

func (b Baron) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (b Baron) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (b Baron) Value() int {
	return 3
}

func (b Baron) Name() string {
	return "Baron"
}

func (b Baron) Description() string {
	return "Compare hands with another player. Player with lower value card is eliminated."
}

func (b Baron) Quantity() int {
	return 2
}
