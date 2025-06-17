package actions

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

func Unmarshal(data []byte) (model.GameAction, error) {
	box, err := coremodel.Unmarshal[model.GameAction](data)
	if err != nil {
		return nil, err
	}
	return box.Value, nil
}
