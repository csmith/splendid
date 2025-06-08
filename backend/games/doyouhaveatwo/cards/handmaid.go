package cards

import (
	"errors"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Handmaid struct{}

func (h Handmaid) Play(game *model.Game, player *model.Player, target *model.Player) error {
	return errors.New("not implemented")
}

func (h Handmaid) CanTarget(game *model.Game, player *model.Player, target *model.Player) bool {
	return false
}

func (h Handmaid) Value() int {
	return 4
}

func (h Handmaid) Name() string {
	return "Handmaid"
}

func (h Handmaid) Description() string {
	return "Player cannot be targeted by other players' cards until their next turn."
}

func (h Handmaid) Quantity() int {
	return 2
}