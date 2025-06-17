package webserver

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/csmith/aca"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

// GameSession represents an active game session
type GameSession struct {
	ID       string
	Engine   *doyouhaveatwo.Engine
	Clients  map[model.PlayerID]*Client
	UpdateCh chan model.GameUpdate
	mu       sync.RWMutex
}

// Client represents a WebSocket connection for a player
type Client struct {
	PlayerID     model.PlayerID
	PlayerName   string
	MessageCh    chan []byte
	DisconnectCh chan struct{}
}

// GameManager manages all active game sessions
type GameManager struct {
	sessions  map[string]*GameSession
	generator *aca.Generator
	logDir    string
	mu        sync.RWMutex
}

func NewGameManager(logDir string) (*GameManager, error) {
	generator, err := aca.NewDefaultGenerator()
	if err != nil {
		return nil, fmt.Errorf("failed to create ID generator: %w", err)
	}

	return &GameManager{
		sessions:  make(map[string]*GameSession),
		generator: generator,
		logDir:    logDir,
	}, nil
}

func (gm *GameManager) CreateGame() (string, error) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	sessionID := gm.generator.Generate()

	// Create update channel for this game session
	updateCh := make(chan model.GameUpdate, 100)

	// Create engine with JSONL logger
	logger := doyouhaveatwo.NewJSONLEventLogger(gm.logDir, sessionID)
	game := model.Game{
		Players:       []*model.Player{},
		Deck:          []coremodel.Redactable[model.Card]{},
		CurrentPlayer: 0,
		Round:         0,
		Phase:         model.PhaseSetup,
		TokensToWin:   4,
	}
	engine := doyouhaveatwo.NewEngine(game, updateCh, logger)

	session := &GameSession{
		ID:       sessionID,
		Engine:   engine,
		Clients:  make(map[model.PlayerID]*Client),
		UpdateCh: updateCh,
	}

	gm.sessions[sessionID] = session

	// Start goroutine to handle game updates
	go gm.handleGameUpdates(session)

	return sessionID, nil
}

func (gm *GameManager) GetSession(sessionID string) (*GameSession, bool) {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	session, exists := gm.sessions[sessionID]
	return session, exists
}

func (gm *GameManager) AddPlayerToSession(sessionID string) (*Client, error) {
	session, exists := gm.GetSession(sessionID)
	if !exists {
		return nil, fmt.Errorf("session not found: %s", sessionID)
	}

	session.mu.Lock()
	defer session.mu.Unlock()

	// Generate player ID and name (same value)
	playerIdentifier := gm.generator.Generate()
	playerID := model.PlayerID(playerIdentifier)

	// Check if player already exists (shouldn't happen with random IDs but just in case)
	if _, exists := session.Clients[playerID]; exists {
		return nil, fmt.Errorf("player already exists in session")
	}

	// Create client
	client := &Client{
		PlayerID:     playerID,
		PlayerName:   playerIdentifier,
		MessageCh:    make(chan []byte, 100),
		DisconnectCh: make(chan struct{}),
	}

	session.Clients[playerID] = client

	// Add player to game if not already added
	existingPlayer := session.Engine.Game.GetPlayer(playerID)
	if existingPlayer == nil {
		// Use the engine's ProcessServerAction to add the player
		addPlayerAction := &actions.AddPlayerAction{
			NewPlayerID:   playerID,
			NewPlayerName: playerIdentifier,
		}

		if err := session.Engine.ProcessServerAction(addPlayerAction); err != nil {
			// Remove client if adding player failed
			delete(session.Clients, playerID)
			return nil, fmt.Errorf("failed to add player to game: %w", err)
		}
	}

	return client, nil
}

func (gm *GameManager) RemovePlayerFromSession(sessionID string, playerID model.PlayerID) {
	session, exists := gm.GetSession(sessionID)
	if !exists {
		return
	}

	session.mu.Lock()
	defer session.mu.Unlock()

	if client, exists := session.Clients[playerID]; exists {
		close(client.DisconnectCh)
		delete(session.Clients, playerID)
	}
}

func (gm *GameManager) handleGameUpdates(session *GameSession) {
	for update := range session.UpdateCh {
		session.mu.RLock()

		// Broadcast update to all connected clients
		for playerID, client := range session.Clients {
			// Create the message wrapper
			gameUpdateMsg := GameUpdateMessage{
				Game:             *update.Game,
				Event:            update.Event,
				AvailableActions: update.AvailableActions,
			}

			msg := WebSocketMessage{
				Type: MessageTypeGameUpdate,
				Data: json.RawMessage{}, // Will be filled after redaction
			}

			// First marshal the game update message
			gameUpdateData, err := json.Marshal(gameUpdateMsg)
			if err != nil {
				slog.Error("Failed to marshal game update message", "error", err, "sessionID", session.ID, "playerID", playerID)
				continue
			}
			msg.Data = gameUpdateData

			// Now redact the entire message for this player
			redactedMsgBytes, err := Redact(msg, playerID)
			if err != nil {
				slog.Error("Failed to redact message for player", "error", err, "sessionID", session.ID, "playerID", playerID)
				continue
			}

			select {
			case client.MessageCh <- redactedMsgBytes:
			case <-client.DisconnectCh:
				// Client disconnected, skip
			default:
				// Channel full, skip to avoid blocking
				slog.Error("Client message channel full, dropping message", "sessionID", session.ID, "playerID", playerID)
			}
		}

		session.mu.RUnlock()
	}
}

func (gm *GameManager) ProcessAction(sessionID string, playerID model.PlayerID, action model.GameAction) error {
	session, exists := gm.GetSession(sessionID)
	if !exists {
		return fmt.Errorf("session not found: %s", sessionID)
	}

	return session.Engine.ProcessAction(playerID, action)
}
