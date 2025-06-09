package webserver

import (
	"encoding/json"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

// Message types for WebSocket communication

type MessageType string

const (
	MessageTypeGameUpdate MessageType = "game_update"
	MessageTypeAction     MessageType = "action"
	MessageTypeError      MessageType = "error"
	MessageTypePlayerID   MessageType = "player_id"
)

// WebSocketMessage is the wrapper for all WebSocket communication
type WebSocketMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

// GameUpdateMessage contains the game state sent to clients
type GameUpdateMessage struct {
	Game             model.Game                                                                     `json:"game"`
	Event            *serialization.Box[model.Event]                                                `json:"event,omitempty"`
	AvailableActions map[model.PlayerID]serialization.Redactable[[]serialization.Box[model.Action]] `json:"available_actions"`
}

// ErrorMessage contains error information
type ErrorMessage struct {
	Message string `json:"message"`
}

// PlayerIDMessage contains the player ID for a connected client
type PlayerIDMessage struct {
	PlayerID string `json:"player_id"`
}

// CreateGameResponse is returned when creating a new game
type CreateGameResponse struct {
	SessionID string `json:"session_id"`
}
