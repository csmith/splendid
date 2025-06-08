package webserver

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

// Message types for WebSocket communication

type MessageType string

const (
	MessageTypeGameUpdate MessageType = "game_update"
	MessageTypeAction     MessageType = "action"
	MessageTypeError      MessageType = "error"
)

// WebSocketMessage is the wrapper for all WebSocket communication
type WebSocketMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

// GameUpdateMessage contains the game state sent to clients
type GameUpdateMessage struct {
	Game             model.Game                                          `json:"game"`
	Event            model.Event                                         `json:"event,omitempty"`
	AvailableActions map[model.PlayerID]model.Redactable[[]model.Action] `json:"available_actions"`
}

// ErrorMessage contains error information
type ErrorMessage struct {
	Message string `json:"message"`
}

// CreateGameResponse is returned when creating a new game
type CreateGameResponse struct {
	SessionID string `json:"session_id"`
}
