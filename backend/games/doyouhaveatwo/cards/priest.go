package cards

import (
	"errors"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Priest struct{}

func (p Priest) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (p Priest) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (p Priest) Value() int {
	return 2
}

func (p Priest) Name() string {
	return "Priest"
}

func (p Priest) Description() string {
	return "Look at another player's hand."
}

func (p Priest) Quantity() int {
	return 2
}