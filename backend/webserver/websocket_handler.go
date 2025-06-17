package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	model2 "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/model"

	"github.com/coder/websocket"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
)

type WebSocketHandler struct {
	gameManager *GameManager
}

func NewWebSocketHandler(gameManager *GameManager) *WebSocketHandler {
	return &WebSocketHandler{
		gameManager: gameManager,
	}
}

func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Extract session ID from URL path
	sessionID := r.PathValue("sessionId")
	if sessionID == "" {
		http.Error(w, "Session ID required", http.StatusBadRequest)
		return
	}

	// Check if session exists
	session, exists := h.gameManager.GetSession(sessionID)
	if !exists {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	// Accept WebSocket connection
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"}, // Allow all origins for now
	})
	if err != nil {
		slog.Error("Failed to accept WebSocket connection", "error", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "Connection closed")

	// Add player to session
	client, err := h.gameManager.AddPlayerToSession(sessionID)
	if err != nil {
		slog.Error("Failed to add player to session", "error", err, "sessionID", sessionID)
		conn.Close(websocket.StatusPolicyViolation, "Failed to join game")
		return
	}

	// Clean up when connection ends
	defer func() {
		h.gameManager.RemovePlayerFromSession(sessionID, client.PlayerID)
		slog.Info("Player disconnected", "sessionID", sessionID, "playerID", client.PlayerID)
	}()

	slog.Info("Player connected", "sessionID", sessionID, "playerID", client.PlayerID, "playerName", client.PlayerName)

	// Send player ID to client
	h.sendPlayerID(conn, client.PlayerID)

	// Send initial game state
	h.sendInitialGameState(conn, session, client.PlayerID)

	// Start goroutines for reading and writing
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Goroutine to write messages to client
	go h.writeMessages(ctx, conn, client)

	// Read messages from client
	h.readMessages(ctx, conn, sessionID, client.PlayerID)
}

func (h *WebSocketHandler) sendInitialGameState(conn *websocket.Conn, session *GameSession, playerID model.PlayerID) {
	// Create a dummy game update with current state
	gameUpdate := model.GameUpdate[model2.Game]{
		Game:             session.Engine.Game,
		Event:            nil, // No specific event for initial state
		AvailableActions: make(map[model.PlayerID]model.Redactable[[]*model.Box[model.Action[model2.Game]]]),
	}

	// For now, we'll create an empty map and let the regular update flow handle actions
	// The engine will send proper updates with actions when events occur

	gameUpdateMsg := GameUpdateMessage[model2.Game]{
		Game:             gameUpdate.Game,
		Event:            gameUpdate.Event,
		AvailableActions: gameUpdate.AvailableActions,
	}

	msg := WebSocketMessage{
		Type: MessageTypeGameUpdate,
	}

	// Marshal the game update message
	gameUpdateData, err := json.Marshal(gameUpdateMsg)
	if err != nil {
		slog.Error("Failed to marshal initial game state", "error", err, "sessionID", session.ID, "playerID", playerID)
		return
	}
	msg.Data = gameUpdateData

	// Redact the entire message for this player
	redactedMsgBytes, err := Redact(msg, playerID)
	if err != nil {
		slog.Error("Failed to redact initial game state", "error", err, "sessionID", session.ID, "playerID", playerID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = conn.Write(ctx, websocket.MessageText, redactedMsgBytes)
	if err != nil {
		slog.Error("Failed to send initial game state", "error", err, "sessionID", session.ID, "playerID", playerID)
	}
}

func (h *WebSocketHandler) writeMessages(ctx context.Context, conn *websocket.Conn, client *Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-client.DisconnectCh:
			return
		case msg := <-client.MessageCh:
			writeCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			err := conn.Write(writeCtx, websocket.MessageText, msg)
			cancel()

			if err != nil {
				slog.Error("Failed to write message to WebSocket", "error", err, "playerID", client.PlayerID)
				return
			}
		}
	}
}

func (h *WebSocketHandler) readMessages(ctx context.Context, conn *websocket.Conn, sessionID string, playerID model.PlayerID) {
	// Set up ping/pong to keep connection alive
	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	// Start a goroutine to handle pings
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-pingTicker.C:
				pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				err := conn.Ping(pingCtx)
				cancel()
				if err != nil {
					slog.Error("Failed to ping client", "error", err, "playerID", playerID)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Remove timeout - let the connection stay open indefinitely with ping/pong keepalive
		_, msgBytes, err := conn.Read(ctx)

		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				slog.Info("WebSocket closed normally", "playerID", playerID)
			} else {
				slog.Error("Failed to read from WebSocket", "error", err, "playerID", playerID)
			}
			return
		}

		// Parse the message
		var msg WebSocketMessage
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			slog.Error("Failed to unmarshal WebSocket message", "error", err, "playerID", playerID)
			h.sendError(conn, "Invalid message format")
			continue
		}

		// Handle different message types
		switch msg.Type {
		case MessageTypeAction:
			h.handleActionMessage(conn, sessionID, playerID, msg.Data)
		default:
			slog.Error("Unknown message type", "type", msg.Type, "playerID", playerID)
			h.sendError(conn, fmt.Sprintf("Unknown message type: %s", msg.Type))
		}
	}
}

func (h *WebSocketHandler) handleActionMessage(conn *websocket.Conn, sessionID string, playerID model.PlayerID, data json.RawMessage) {
	action, err := actions.Unmarshal(data)
	if err != nil {
		slog.Error("Failed to unmarshal action", "error", err, "playerID", playerID)
		h.sendError(conn, "Invalid action format")
		return
	}

	// Process the action
	if err := h.gameManager.ProcessAction(sessionID, playerID, action); err != nil {
		slog.Error("Failed to process action", "error", err, "sessionID", sessionID, "playerID", playerID)
		h.sendError(conn, fmt.Sprintf("Action failed: %s", err.Error()))
		return
	}
}

func (h *WebSocketHandler) sendPlayerID(conn *websocket.Conn, playerID model.PlayerID) {
	playerIDData, err := json.Marshal(string(playerID))
	if err != nil {
		slog.Error("Failed to marshal player ID", "error", err)
		return
	}

	msg := WebSocketMessage{
		Type: MessageTypePlayerID,
		Data: playerIDData,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		slog.Error("Failed to marshal player ID WebSocket message", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = conn.Write(ctx, websocket.MessageText, msgBytes)
	if err != nil {
		slog.Error("Failed to send player ID message", "error", err)
	}
}

func (h *WebSocketHandler) sendError(conn *websocket.Conn, message string) {
	errorMsg := ErrorMessage{
		Message: message,
	}

	msg := WebSocketMessage{
		Type: MessageTypeError,
	}

	errorData, err := json.Marshal(errorMsg)
	if err != nil {
		slog.Error("Failed to marshal error message", "error", err)
		return
	}
	msg.Data = errorData

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		slog.Error("Failed to marshal error WebSocket message", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = conn.Write(ctx, websocket.MessageText, msgBytes)
	if err != nil {
		slog.Error("Failed to send error message", "error", err)
	}
}
