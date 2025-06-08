package cards

import (
	"errors"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Guard struct{}

func (g Guard) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (g Guard) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (g Guard) Value() int {
	return 1
}

func (g Guard) Name() string {
	return "Guard"
}

func (g Guard) Description() string {
	return "Guess another player's card. If correct, that player is eliminated."
}

func (g Guard) Quantity() int {
	return 5
}
