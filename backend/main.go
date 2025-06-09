package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/csmith/hotsource"
	"github.com/csmith/splendid/frontend"

	"github.com/csmith/splendid/backend/webserver"
)

func main() {
	// Parse command line flags
	logDir := flag.String("data", "data", "Directory to store game logs")
	flag.Parse()

	// Set up structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Create game manager
	gameManager, err := webserver.NewGameManager(*logDir)
	if err != nil {
		slog.Error("Failed to create game manager", "error", err)
		os.Exit(1)
	}

	// Create handlers
	httpHandler := webserver.NewHTTPHandler(gameManager)
	wsHandler := webserver.NewWebSocketHandler(gameManager)

	// Set up routes
	mux := http.NewServeMux()

	mux.Handle("/", hotsource.Maybe(http.FileServerFS(frontend.FS), "frontend", true))

	// HTTP endpoints
	mux.HandleFunc("POST /api/games", httpHandler.CreateGame)

	// WebSocket endpoints
	mux.HandleFunc("GET /api/games/{sessionId}", wsHandler.HandleWebSocket)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	slog.Info("Starting server", "address", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
