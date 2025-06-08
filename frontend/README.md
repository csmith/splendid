# Frontend for Do You Have A Two

A simple HTML+JavaScript frontend for the "Do You Have A Two" card game.

## Usage

1. **Start the backend server:**
   ```bash
   cd backend
   go run main.go
   ```
   The server will start on port 8080 by default.

2. **Open the frontend:**
   Open `index.html` in your web browser. You can either:
   - Double-click the file to open it locally
   - Serve it via a simple HTTP server (recommended for WebSocket compatibility)

3. **Play the game:**
   - Click "Create New Game" to start a new session
   - Copy the session ID to share with other players
   - Or enter an existing session ID to join a game
   - Click "Join Game" to connect

## Features

- **Session Management:** Create new games or join existing ones
- **Real-time Gameplay:** WebSocket connection for live updates
- **Multi-step Actions:** Supports complex card interactions like Guard targeting
- **Event Log:** Complete history of game events
- **Card Information:** Click cards to see their effects
- **Player Status:** Visual indicators for protection, elimination, current player
- **Responsive Design:** Works on desktop and mobile devices

## Game Flow

1. **Setup:** Add players (2-4) and start the game
2. **Play:** Current player draws a card, then plays one
3. **Actions:** Use buttons to play cards and target other players
4. **Win:** First player to reach the token threshold wins!

## WebSocket API

The frontend communicates with the backend using WebSocket messages:

- **Outgoing:** `{ "type": "action", "data": { ... } }`
- **Incoming:** `{ "type": "game_update", "data": { ... } }`
- **Errors:** `{ "type": "error", "data": { "message": "..." } }`

## Card Effects

- **Guard (1):** Guess another player's card
- **Priest (2):** Look at another player's hand
- **Baron (3):** Compare hands with another player
- **Handmaid (4):** Protection until next turn
- **Prince (5):** Force a player to discard and draw
- **King (6):** Trade hands with another player
- **Countess (7):** Must discard if holding King or Prince
- **Princess (8):** If discarded, player is eliminated