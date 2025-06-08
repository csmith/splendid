package actions

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

func Unmarshal(data []byte) (model.Action, error) {
	// First, unmarshal just the type field
	var actionType struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &actionType); err != nil {
		return nil, fmt.Errorf("failed to unmarshal action type: %w", err)
	}

	// Based on the type, declare the correct concrete type
	var action model.Action
	switch actionType.Type {
	case "add_player":
		action = &AddPlayerAction{}
	case "draw_card":
		action = &DrawCardAction{}
	case "start_game":
		action = &StartGameAction{}
	case "play_guard":
		action = &PlayCardGuardAction{}
	case "play_card_no_target":
		action = &PlayCardNoTargetAction{}
	case "play_card_target_any":
		action = &PlayCardTargetAnyAction{}
	case "play_card_target_others":
		action = &PlayCardTargetOthersAction{}
	case "discard_card":
		action = &DiscardCardAction{}
	default:
		return nil, fmt.Errorf("unknown action type: %s", actionType.Type)
	}

	// Unmarshal the full JSON into the concrete type
	if err := json.Unmarshal(data, action); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s: %w", actionType.Type, err)
	}

	return action, nil
}
