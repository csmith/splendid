package cards

import (
	"errors"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Countess struct{}

func (c Countess) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (c Countess) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (c Countess) Value() int {
	return 7
}

func (c Countess) Name() string {
	return "Countess"
}

func (c Countess) Description() string {
	return "No special effect, but must be discarded if holding King or Prince."
}

func (c Countess) Quantity() int {
	return 1
}