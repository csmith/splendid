package actions

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

func Unmarshal(data []byte) (model.Action, error) {
	box, err := serialization.Unmarshal[model.Action](data)
	if err != nil {
		return nil, err
	}
	return box.Value, nil
}
