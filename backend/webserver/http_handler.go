package webserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HTTPHandler struct {
	gameManager *GameManager
}

func NewHTTPHandler(gameManager *GameManager) *HTTPHandler {
	return &HTTPHandler{
		gameManager: gameManager,
	}
}

func (h *HTTPHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID, err := h.gameManager.CreateGame()
	if err != nil {
		slog.Error("Failed to create game", "error", err)
		http.Error(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	response := CreateGameResponse{
		SessionID: sessionID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode create game response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	slog.Info("Game created", "sessionID", sessionID)
}
