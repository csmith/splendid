// Global state
let gameState = null;
let ws = null;
let currentPlayerID = null;
let sessionID = null;
let availableActions = [];

// API endpoints
const API_BASE = window.location.origin.replace(/^http/, 'ws');
const HTTP_BASE = window.location.origin;

// Card information for tooltips
const CARD_INFO = {
    'Guard': { value: 1, description: 'Guess another player\'s card. If correct, that player is eliminated.', quantity: 5 },
    'Priest': { value: 2, description: 'Look at another player\'s hand.', quantity: 2 },
    'Baron': { value: 3, description: 'Compare hands with another player. Player with lower value is eliminated.', quantity: 2 },
    'Handmaid': { value: 4, description: 'Cannot be targeted by other players until your next turn.', quantity: 2 },
    'Prince': { value: 5, description: 'Target player discards their hand and draws a new card.', quantity: 2 },
    'King': { value: 6, description: 'Trade hands with another player.', quantity: 1 },
    'Countess': { value: 7, description: 'Must be discarded if holding King or Prince.', quantity: 1 },
    'Princess': { value: 8, description: 'If discarded, player is eliminated.', quantity: 1 }
};

// Function to convert quantity to braille dots
function getBrailleDots(quantity) {
    const brailleNumbers = ['⠀', '⠈', '⠉', '⠙', '⠛', '⠻'];
    return brailleNumbers[quantity] || '⠻';
}

// Initialize the application
document.addEventListener('DOMContentLoaded', function() {
    updateConnectionStatus('Not connected');
});

// Session management functions
async function createGame() {
    try {
        updateConnectionStatus('Creating game...');
        const response = await fetch(`${HTTP_BASE}/api/games`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        document.getElementById('session-id').value = data.session_id;
        updateConnectionStatus('Game created! Click "Join Game" to connect.');
        logEvent('Game created', `Session ID: ${data.session_id}`);
    } catch (error) {
        console.error('Error creating game:', error);
        showError('Failed to create game: ' + error.message);
        updateConnectionStatus('Failed to create game');
    }
}

function joinGame() {
    const sessionIdInput = document.getElementById('session-id');
    const inputSessionID = sessionIdInput.value.trim();
    
    if (!inputSessionID) {
        showError('Please enter a session ID');
        return;
    }

    sessionID = inputSessionID;
    connectWebSocket(sessionID);
}

function connectWebSocket(sessionID) {
    try {
        updateConnectionStatus('Connecting...');
        
        // Close existing connection if any
        if (ws) {
            ws.close();
        }

        // Connect to WebSocket
        const wsUrl = `${API_BASE}/api/games/${sessionID}`;
        ws = new WebSocket(wsUrl);

        ws.onopen = function(event) {
            updateConnectionStatus('Connected');
            showGameSection();
            logEvent('Connected', 'Successfully joined game');
        };

        ws.onmessage = function(event) {
            try {
                const message = JSON.parse(event.data);
                handleWebSocketMessage(message);
            } catch (error) {
                console.error('Error parsing message:', error);
                showError('Error parsing server message');
            }
        };

        ws.onclose = function(event) {
            updateConnectionStatus('Disconnected');
            hideGameSection();
            if (event.code !== 1000) { // Not a normal closure
                showError('Connection lost: ' + (event.reason || 'Unknown reason'));
            }
        };

        ws.onerror = function(error) {
            console.error('WebSocket error:', error);
            showError('Connection error');
            updateConnectionStatus('Connection error');
        };

    } catch (error) {
        console.error('Error connecting:', error);
        showError('Failed to connect: ' + error.message);
        updateConnectionStatus('Failed to connect');
    }
}

function handleWebSocketMessage(message) {
    switch (message.type) {
        case 'player_id':
            currentPlayerID = message.data;
            logEvent('Player ID received', currentPlayerID);
            break;
        case 'game_update':
            handleGameUpdate(message.data);
            break;
        case 'error':
            showError(message.data.message);
            break;
        default:
            console.warn('Unknown message type:', message.type);
    }
}

function handleGameUpdate(gameUpdateData) {
    gameState = gameUpdateData.game;
    
    // Update available actions - find actions for our player ID
    availableActions = [];
    
    if (gameUpdateData.available_actions && currentPlayerID) {
        const playerActions = gameUpdateData.available_actions[currentPlayerID];
        if (playerActions && playerActions !== "REDACTED" && Array.isArray(playerActions)) {
            availableActions = playerActions;
        }
    }

    // Log the event if there is one
    if (gameUpdateData.event) {
        logGameEvent(gameUpdateData.event);
    }

    // Update UI
    updateGameUI();
}

function getCurrentPlayerID() {
    // currentPlayerID is now set when we receive the player_id message
    return currentPlayerID;
}

function updateGameUI() {
    updateGameStatus();
    updatePlayersList();
    updateDeckDisplay();
    updateYourHand();
    updateActionsList();
}

function updateGameStatus() {
    if (!gameState) return;

    document.getElementById('round-info').textContent = `Round: ${gameState.round || 0}`;
    document.getElementById('phase-info').textContent = `Phase: ${gameState.phase || 'Unknown'}`;
    document.getElementById('tokens-to-win-info').textContent = `Tokens to win: ${gameState.tokens_to_win || 'Unknown'}`;
    
    const currentPlayerIndex = gameState.current_player || 0;
    const currentPlayer = gameState.players && gameState.players[currentPlayerIndex];
    document.getElementById('current-player-info').textContent = `Current Player: ${currentPlayer ? currentPlayer.name : 'Unknown'}`;
}

function updatePlayersList() {
    const playersContainer = document.getElementById('players-list');
    playersContainer.innerHTML = '';

    if (!gameState || !gameState.players) return;

    gameState.players.forEach((player, index) => {
        const playerDiv = document.createElement('div');
        playerDiv.className = 'player';
        
        // Add classes for player state
        if (player.is_out) playerDiv.classList.add('eliminated');
        if (player.is_protected) playerDiv.classList.add('protected');
        if (index === gameState.current_player) playerDiv.classList.add('current');
        if (player.id === currentPlayerID) playerDiv.classList.add('you');

        const handSize = player.hand ? player.hand.length : 0;
        const discardPile = player.discard_pile || [];

        playerDiv.innerHTML = `
            <div class="player-name">${player.name} ${player.id === currentPlayerID ? '(You)' : ''}</div>
            <div class="player-stats">
                <span class="tokens">Tokens: ${player.token_count || 0}</span>
                <span class="hand-size">Cards: ${handSize}</span>
            </div>
            <div class="player-status">
                ${player.is_out ? '<span class="status eliminated">Eliminated</span>' : ''}
                ${player.is_protected ? '<span class="status protected">Protected</span>' : ''}
            </div>
            ${discardPile.length > 0 ? `<div class="discard-pile">
                <div class="discard-title">Played cards:</div>
                <div class="discard-cards"></div>
            </div>` : ''}
        `;

        // Add discard pile cards if any exist
        if (discardPile.length > 0) {
            const discardCardsContainer = playerDiv.querySelector('.discard-cards');
            discardPile.forEach(cardName => {
                const cardDiv = document.createElement('div');
                cardDiv.className = 'card discard-card';
                
                if (cardName === "REDACTED") {
                    cardDiv.classList.add('redacted');
                    cardDiv.innerHTML = `
                        <div class="card-value">?</div>
                        <div class="card-name">Hidden</div>
                    `;
                } else if (typeof cardName === 'string') {
                    const cardInfo = CARD_INFO[cardName];
                    if (cardInfo) {
                        cardDiv.onclick = () => showCardInfo(cardName);
                        cardDiv.innerHTML = `
                            <div class="card-quantity">${getBrailleDots(cardInfo.quantity)}</div>
                            <div class="card-value">${cardInfo.value}</div>
                            <div class="card-name">${cardName}</div>
                        `;
                    }
                }
                
                discardCardsContainer.appendChild(cardDiv);
            });
        }

        playersContainer.appendChild(playerDiv);
    });
}

function updateDeckDisplay() {
    const deckContainer = document.getElementById('deck-display');
    deckContainer.innerHTML = '';

    if (!gameState || !gameState.deck) {
        // Show empty deck outline
        const emptyDeck = document.createElement('div');
        emptyDeck.className = 'deck-empty';
        emptyDeck.textContent = 'Empty Deck';
        deckContainer.appendChild(emptyDeck);
        return;
    }

    const deckSize = gameState.deck.length;
    
    if (deckSize === 0) {
        // Show empty deck outline
        const emptyDeck = document.createElement('div');
        emptyDeck.className = 'deck-empty';
        emptyDeck.textContent = 'Empty Deck';
        deckContainer.appendChild(emptyDeck);
        return;
    }

    // Show staggered deck cards
    // We'll show up to 5 cards for visual effect, but the real count on the top card
    const cardsToShow = Math.min(5, deckSize);
    
    for (let i = 0; i < cardsToShow; i++) {
        const deckCard = document.createElement('div');
        deckCard.className = 'deck-card';
        
        // Add card back design
        const cardBack = document.createElement('div');
        cardBack.className = 'deck-card-back';
        cardBack.innerHTML = '?';
        deckCard.appendChild(cardBack);
        
        // Add count only to the top (last) card
        if (i === cardsToShow - 1) {
            const countBadge = document.createElement('div');
            countBadge.className = 'deck-count';
            countBadge.textContent = deckSize.toString();
            deckCard.appendChild(countBadge);
        }
        
        deckContainer.appendChild(deckCard);
    }
}

function updateYourHand() {
    const handContainer = document.getElementById('your-hand');
    handContainer.innerHTML = '';

    if (!gameState || !currentPlayerID || !gameState.players) return;

    const currentPlayer = gameState.players.find(p => p.id === currentPlayerID);
    if (!currentPlayer || !currentPlayer.hand) return;

    currentPlayer.hand.forEach(handCard => {
        const cardDiv = document.createElement('div');
        cardDiv.className = 'card';
        
        if (handCard === "REDACTED") {
            cardDiv.classList.add('redacted');
            cardDiv.innerHTML = `
                <div class="card-value">?</div>
                <div class="card-name">Hidden</div>
            `;
        } else if (typeof handCard === 'string') {
            // Card name as string
            const cardInfo = CARD_INFO[handCard];
            if (cardInfo) {
                cardDiv.onclick = () => showCardInfo(handCard);
                cardDiv.innerHTML = `
                    <div class="card-quantity">${getBrailleDots(cardInfo.quantity)}</div>
                    <div class="card-value">${cardInfo.value}</div>
                    <div class="card-name">${handCard}</div>
                `;
            }
        }

        handContainer.appendChild(cardDiv);
    });
}

function updateActionsList() {
    const actionsContainer = document.getElementById('actions-list');
    actionsContainer.innerHTML = '';

    if (!availableActions || availableActions.length === 0) {
        actionsContainer.innerHTML = '<p>No actions available</p>';
        return;
    }

    availableActions.forEach(action => {
        const button = document.createElement('button');
        button.className = 'action-btn';
        button.textContent = formatActionText(action);
        button.onclick = () => sendAction(action);
        actionsContainer.appendChild(button);
    });
}

function formatActionText(action) {
    // Extract action type (remove dyhat:a: prefix if present)
    const actionType = action.type.startsWith('dyhat:a:') ? action.type.substring(8) : action.type;
    
    switch (actionType) {
        case 'add_player':
            return 'Add Player';
        case 'start_game':
            return 'Start Game';
        case 'draw_card':
            return 'Draw Card';
        case 'play_card_with_guess':
            // Guard card actions
            if (!action.target_player && !action.guessed_rank) {
                return `Play ${action.card_name || 'Guard'}`;
            } else if (!action.guessed_rank) {
                return `${action.card_name || 'Guard'}: target ${getPlayerName(action.target_player)}`;
            } else {
                return `${action.card_name || 'Guard'}: guess ${action.guessed_rank} (${getCardName(action.guessed_rank)})`;
            }
        case 'play_card_target_others':
            // Baron, Priest, King cards
            if (!action.target_player) {
                return `Play ${action.card_name}`;
            } else {
                return `${action.card_name}: target ${getPlayerName(action.target_player)}`;
            }
        case 'play_card_target_any':
            // Prince card
            if (!action.target_player) {
                return `Play ${action.card_name}`;
            } else {
                return `${action.card_name}: target ${getPlayerName(action.target_player)}`;
            }
        case 'play_card_no_target':
            // Handmaid, Countess, Princess cards
            return `Play ${action.card_name}`;
        case 'discard_card':
            return `Play ${action.card_name} (no effect)`;
        default:
            // Convert snake_case to Title Case for unknown action types
            return actionType
                .split('_')
                .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                .join(' ');
    }
}

function getPlayerName(playerID) {
    if (!gameState || !gameState.players) return playerID;
    const player = gameState.players.find(p => p.id === playerID);
    return player ? player.name : playerID;
}

function getCardName(rank) {
    const cardEntry = Object.entries(CARD_INFO).find(([name, info]) => info.value === rank);
    return cardEntry ? cardEntry[0] : `Card ${rank}`;
}

function getCardNameFromAction(actionType) {
    // Convert action type like "play_priest" to "Priest"
    if (actionType.startsWith('play_')) {
        const cardName = actionType.substring(5); // Remove "play_"
        return cardName.charAt(0).toUpperCase() + cardName.slice(1);
    }
    return actionType;
}

function sendAction(action) {
    if (!ws || ws.readyState !== WebSocket.OPEN) {
        showError('Not connected to game');
        return;
    }

    const message = {
        type: 'action',
        data: action
    };

    try {
        ws.send(JSON.stringify(message));
        logEvent('Action sent', formatActionText(action));
    } catch (error) {
        console.error('Error sending action:', error);
        showError('Failed to send action');
    }
}

function logGameEvent(event) {
    if (!event) return;
    
    let eventText;
    
    // Extract event type (remove dyhat:e: prefix if present)
    const eventType = event.type.startsWith('dyhat:e:') ? event.type.substring(8) : event.type;
    
    // Add event-specific details
    switch (eventType) {
        case 'player_added':
            eventText = `Player ${getPlayerName(event.player)} joined the game`;
            break;
        case 'card_dealt':
            eventText = `Card dealt to ${getPlayerName(event.to_player)}`;
            break;
        case 'card_discarded':
            if (event.discarded_card) {
                eventText = `${getPlayerName(event.player)} discarded ${event.discarded_card}`;
            } else {
                eventText = `${getPlayerName(event.player)} discarded a card`;
            }
            break;
        case 'player_eliminated':
            eventText = `${getPlayerName(event.player)} was eliminated`;
            break;
        case 'player_protection_granted':
            eventText = `${getPlayerName(event.player)} is now protected`;
            break;
        case 'player_protection_cleared':
            eventText = `${getPlayerName(event.player)} is no longer protected`;
            break;
        case 'round_updated':
            eventText = `Round ${event.round || '?'} started`;
            break;
        case 'phase_updated':
            eventText = `Phase changed to ${event.phase || 'unknown'}`;
            break;
        case 'current_player_updated':
            if (gameState && gameState.players && event.current_player !== undefined) {
                const player = gameState.players[event.current_player];
                eventText = `It's now ${player ? player.name : 'unknown'}'s turn`;
            } else {
                eventText = 'Current player changed';
            }
            break;
        case 'player_token_awarded':
            eventText = `${getPlayerName(event.player)} won a token!`;
            break;
        case 'winner_declared':
            eventText = `🎉 ${getPlayerName(event.winner)} wins the game! 🎉`;
            break;
        case 'deck_updated':
            eventText = `Deck updated (${event.deck ? event.deck.length : 0} cards remaining)`;
            break;
        case 'card_removed':
            if (event.removed_card && event.removed_card !== "REDACTED") {
                eventText = `${event.removed_card} was removed from the deck`;
            } else {
                eventText = `A card was removed from the deck`;
            }
            break;
        case 'player_restored':
            eventText = `${getPlayerName(event.player)} was restored to the game`;
            break;
        case 'player_discard_pile_cleared':
            eventText = `${getPlayerName(event.player)}'s discard pile was cleared`;
            break;
        case 'hand_revealed':
            const sourcePlayer = getPlayerName(event.source_player);
            const targetPlayers = event.target_players && event.target_players.length > 0 
                ? event.target_players.map(getPlayerName).join(', ') 
                : 'someone';
            if (event.revealed_card && event.revealed_card !== "REDACTED") {
                eventText = `${sourcePlayer}'s ${event.revealed_card} was revealed to ${targetPlayers}`;
            } else {
                eventText = `${sourcePlayer}'s card was revealed to ${targetPlayers}`;
            }
            break;
        case 'hands_swapped':
            eventText = `${getPlayerName(event.player1)} and ${getPlayerName(event.player2)} swapped hands`;
            break;
        case 'last_round_winners_updated':
            eventText = `Last round winners updated`;
            break;
        case 'player_action_completed':
            eventText = `${getPlayerName(event.player)} completed their action`;
            break;
        case 'player_action_started':
            eventText = `${getPlayerName(event.player)} started their action`;
            break;
        case 'player_action_updated':
            eventText = `${getPlayerName(event.player)} updated their action`;
            break;
        case 'player_hand_cleared':
            eventText = `${getPlayerName(event.player)}'s hand was cleared`;
            break;
        case 'removed_card_dealt':
            eventText = `The removed card was dealt to ${getPlayerName(event.to_player)}`;
            break;
        default:
            // For unknown event types, just show the type and any available info
            console.log('Unknown event type:', event);
            eventText = JSON.stringify(event);
    }
    
    logEvent('Game Event', eventText);
}

function logEvent(type, message) {
    const eventLog = document.getElementById('event-log');
    const timestamp = new Date().toLocaleTimeString();
    
    const eventDiv = document.createElement('div');
    eventDiv.className = 'event';
    eventDiv.innerHTML = `
        <span class="event-time">${timestamp}</span>
        <span class="event-type">${type}:</span>
        <span class="event-message">${message}</span>
    `;
    
    eventLog.appendChild(eventDiv);
    eventLog.scrollTop = eventLog.scrollHeight;
}

function showError(message) {
    const errorContainer = document.getElementById('error-messages');
    const errorDiv = document.createElement('div');
    errorDiv.className = 'error';
    errorDiv.textContent = message;
    
    // Add close button
    const closeBtn = document.createElement('span');
    closeBtn.className = 'close-error';
    closeBtn.textContent = '×';
    closeBtn.onclick = () => errorDiv.remove();
    errorDiv.appendChild(closeBtn);
    
    errorContainer.appendChild(errorDiv);
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
        if (errorDiv.parentNode) {
            errorDiv.remove();
        }
    }, 5000);
}

function updateConnectionStatus(status) {
    document.getElementById('connection-status').textContent = status;
}

function showGameSection() {
    document.getElementById('session-section').style.display = 'none';
    document.getElementById('game-section').style.display = 'block';
}

function hideGameSection() {
    document.getElementById('session-section').style.display = 'block';
    document.getElementById('game-section').style.display = 'none';
}

function showCardInfo(cardName) {
    const cardInfo = CARD_INFO[cardName];
    if (!cardInfo) return;
    
    document.getElementById('card-info-title').textContent = `${cardName} (${cardInfo.value})`;
    document.getElementById('card-info-description').textContent = cardInfo.description;
    document.getElementById('card-info-modal').style.display = 'block';
}

function closeCardInfo() {
    document.getElementById('card-info-modal').style.display = 'none';
}

// Close modal when clicking outside
window.onclick = function(event) {
    const modal = document.getElementById('card-info-modal');
    if (event.target === modal) {
        closeCardInfo();
    }
}