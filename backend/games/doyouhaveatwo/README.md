# Do You Have A Two — Game Specification

A Love Letter-inspired card game.

## Game Overview

Players compete to deliver their letter to the princess by using intermediaries in the palace. Each round, players play
cards with different powers to eliminate opponents or gain information. The last player standing wins the round and
gains a token. First player to collect enough tokens wins the game.

## Game Components

### Cards

#### Base variant
- **Guard (1) x5**: guess another player's card
- **Priest (2) x2**: look at another player's hand
- **Baron (3) x2**: compare hands with another player (lower value is eliminated)
- **Handmaid (4) x2**: protection until next turn
- **Prince (5) x2**: force a player to discard and draw
- **King (6) x1**: trade hands with another player
- **Countess (7) x1**: must discard if holding King or Prince
- **Princess (8) x1**: if discarded, player is eliminated

### Players
- 2-4 players supported
- Each player starts with 1 card
- Players collect tokens for winning rounds

## Architecture

### Package Structure

The codebase is organized into separate packages for clean separation of concerns:

- **`model/`** - Core types and data structures (Game, Player, Card, Event, etc.)
- **`events/`** - Concrete event types (StartRoundEvent, DrawCardEvent, etc.)
- **Main package** - Game engine and orchestration logic

### Core Types

#### Game (model/game.go)
```go
type Game struct {
    Players       []*Player
    Deck          []Redactable[Card]
    RemovedCard   Redactable[Card]    // Face-down card removed each round
    CurrentPlayer int
    Round         int
    Phase         GamePhase
    TokensToWin   int
}
```

#### Player (model/player.go)
```go
type Player struct {
    ID          PlayerID
    Name        string
    Hand        []Redactable[Card]
    DiscardPile []Card      // Player's own discard pile (visible to all)
    TokenCount  int
    IsOut       bool        // Eliminated this round
    IsProtected bool        // Handmaid protection
    Position    int         // Seating order
}
```

#### Card Type (model/card.go)
```go
type Card struct {
    value       int
    name        string
    description string
    quantity    int
}

var (
    CardGuard    = Card{value: 1, name: "Guard", ...}
    CardPriest   = Card{value: 2, name: "Priest", ...}
    // ... other cards
)

var CardTypes = []Card{
    CardGuard, CardPriest, CardBaron, CardHandmaid,
    CardPrince, CardKing, CardCountess, CardPrincess,
}
```

**Card Data**: All cards are defined as simple struct instances in `model/card.go`:
- `CardGuard` - Guard card (value 1, quantity 5)
- `CardPriest` - Priest card (value 2, quantity 2) 
- `CardBaron` - Baron card (value 3, quantity 2)
- `CardHandmaid` - Handmaid card (value 4, quantity 2)
- `CardPrince` - Prince card (value 5, quantity 2)
- `CardKing` - King card (value 6, quantity 1)
- `CardCountess` - Countess card (value 7, quantity 1)
- `CardPrincess` - Princess card (value 8, quantity 1)

The `CardTypes` slice contains all available card types for easy iteration during deck creation.

**Card Behavior**: Card logic is implemented as individual event types in the `events/` package rather than as methods on card types. This separation allows for better event sourcing and replay capabilities.

#### Redactable Type (model/redactable.go)
```go
type Redactable[T any] struct {
    Value     T
    VisibleTo map[PlayerID]bool
}
```

### Game Flow

#### Setup Phase
1. **Deck Preparation**: Create deck with all 16 cards (5 Guards, 2 Priests, 2 Barons, 2 Handmaids, 2 Princes, 1 King, 1 Countess, 1 Princess)
2. **Shuffle**: Thoroughly randomize the deck
3. **Remove Card**: Take top card and set aside face-down (this card is out of play for the round)
4. **Deal Initial Cards**: Each player receives 1 card face-down
5. **Determine First Player**: Randomly select starting player or use previous round's winner
6. **Initialize State**: Reset all player protection flags, clear individual discard piles, set round counter

#### Play Phase
Each turn consists of the following steps:

1. **Clear Protection**: Remove Handmaid protection from current player (if applicable)
2. **Draw Card**: Current player draws 1 card from deck (now has 2 cards)
3. **Check Forced Plays**: If player holds Countess + (King or Prince), must play Countess
4. **Card Selection**: Player chooses which of their 2 cards to play
5. **Target Selection**: If card requires a target, player selects valid target
6. **Card Resolution**: Execute card effect, update game state
7. **Discard**: Played card goes to player's individual discard pile
8. **Round End Check**: If only one player remains or deck is empty, round ends
9. **Next Player**: Advance to next active player

#### Round End Phase
A round ends when either:
- Only one player remains (others eliminated)
- Deck is empty (all cards drawn)

**Winner Determination**:
1. If one player remains: That player wins the round
2. If multiple players remain: Player with highest-value card wins
3. If tied for highest: All tied players win a token

**Cleanup**:
1. Award token(s) to round winner(s)
2. Reset all player states (IsOut = false, IsProtected = false)
3. Clear hands and individual discard piles
4. Check for game end condition
5. If game continues, start new round with winner(s) going first

#### Game End Phase
**Win Condition**: First player(s) to reach the token threshold wins the game.

**Token Thresholds by Player Count**:
- 2 players: 7 tokens
- 3 players: 5 tokens  
- 4 players: 4 tokens

**Final Resolution**:
1. Declare winner(s)
2. Calculate final scores/statistics
3. Log game completion
4. Clean up game state

**Tie Handling**: Multiple players can win simultaneously if they reach the threshold in the same round.

### Card Effects

#### Guard (Value: 1)
**Effect**: Guess another player's card. If correct, that player is eliminated.

**Targeting**: Must target another player who is not protected.

**Edge Cases**:
- Cannot guess "Guard" (invalid guess)
- If no valid targets exist (all other players protected or eliminated), card has no effect
- Must still choose a target even if guess will be wrong

#### Priest (Value: 2)
**Effect**: Look at another player's hand.

**Targeting**: Must target another player who is not protected.

**Edge Cases**:
- If no valid targets exist, card has no effect but is still played

#### Baron (Value: 3)
**Effect**: Compare hands with another player. Player with lower value card is eliminated. If tied, nothing happens.

**Targeting**: Must target another player who is not protected.

**Edge Cases**:
- If no valid targets exist, card has no effect
- Ties result in no elimination
- Both players' cards remain in their hands (not discarded)

#### Handmaid (Value: 4)
**Effect**: Player cannot be targeted by other players' cards until their next turn.

**Targeting**: No target required (self-effect).

**Edge Cases**:
- Protection lasts until the start of the player's next turn
- Protection is removed when player's next turn begins, not when they play a card

#### Prince (Value: 5)
**Effect**: Target player discards their hand and draws a new card. Can target self.

**Targeting**: Must target any player (including self) who is not protected.

**Edge Cases**:
- If no valid targets exist (all other players protected), must target self
- If deck is empty, target draws the removed card (face-down card becomes available)
- If target discards Princess, they are eliminated
- If held with Countess, Countess must be played instead

#### King (Value: 6)
**Effect**: Trade hands with another player.

**Targeting**: Must target another player who is not protected.

**Edge Cases**:
- If no valid targets exist, card has no effect
- If held with Countess, Countess must be played instead
- Trade is simultaneous — both players get each other's card

#### Countess (Value: 7)
**Effect**: No special effect, but must be discarded if holding King or Prince.

**Targeting**: No target required.

**Edge Cases**:
- Forced discard overrides normal card choice
- Must be played even if it's strategically disadvantageous
- May be manually played, even if not holding a King or Prince.

#### Princess (Value: 8)
**Effect**: If discarded (played or forced to discard), player is immediately eliminated.

**Targeting**: No target required, but playing it eliminates yourself.

**Edge Cases**:
- Never voluntarily played (would eliminate self)
- If forced to discard by Prince, player is eliminated
- Highest value card, so good for end-of-round if still in hand

### Game Events & Logging

Game events are tracked separately from the core Game state to avoid bloating client updates. The server maintains a complete GameLog for:

- Game replay
- Debugging
- Audit trails
- Undo functionality (if needed)

Each Event is sent to clients as part of GameUpdate but stored separately on the server.

## API Design

### Client-Server Architecture

The game server maintains all game logic and state while clients only handle presentation. This keeps clients simple and prevents cheating.

### Communication Types

#### GameUpdate (model/action.go)
What the server sends to each player after every action:
```go
type GameUpdate struct {
    Game             Game                            // Current redacted game state
    Event            Event                           // What just happened (nil for initial state)
    AvailableActions map[PlayerID]Redactable[[]Action] // Actions available to each player
}
```

#### Action (model/action.go)
Structured actions that UIs can handle generically:
```go
type Action struct {
    Type  string      // "play_card", "target_player", "select_guess", etc.
    Value interface{} // Card name, player ID, guess value, etc.
    Label string      // Human-readable text for UI
}
```

#### Event Interface (model/event.go)
Records what happened for event log and client updates:
```go
type Event interface {
    Type() EventType
    PlayerID() *PlayerID
    Apply(g *Game) error
}
```

**Event Implementations**: Each event type is implemented in `events/` package:
- `events/start_round.go` - StartRoundEvent with EventStartRound constant
- `events/draw_card.go` - DrawCardEvent with EventDrawCard constant

```go
// Example concrete event types:
type StartRoundEvent struct{}
type DrawCardEvent struct {
    Player PlayerID
}
type PlayCardEvent struct {
    Player PlayerID
    Card   Card
    Target *PlayerID
}
```

Each event implements the Event interface with pointer receivers for the `Apply(*Game) error` method, providing type safety and eliminating the need for `interface{}` parameters.

**Secret State Revelation**: Events should expose secret information through `Redactable` fields that are populated during the `Apply` function. This allows clients to receive proper logging information while maintaining the visibility system. For example, `DrawCardEvent.CardDrawn` contains the card that was drawn, properly redacted based on player visibility rules.

### State Redaction

Players receive redacted game state based on visibility rules:

**Public Information** (visible to all players):
- All player names, positions, token counts
- All player protection/elimination status
- Hand sizes (number of cards, not actual cards)
- All discard piles (cards visible)
- Current player, round, phase
- Deck size

**Private Information** (visible only to card owner):
- Actual cards in hand

### Action Flow

Actions follow a multi-step approach to handle targeting and choices:

1. **Available Actions**: Server sends GameUpdate with possible actions
   ```go
   []Action{
       {Type: "play_card", Value: "guard", Label: "Play Guard"},
       {Type: "play_card", Value: "priest", Label: "Play Priest"},
   }
   ```

2. **Action Selection**: Player chooses an action
   ```go
   Action{Type: "play_card", Value: "guard"}
   ```

3. **Sub-Actions** (if needed): Server responds with follow-up choices
   ```go
   []Action{
       {Type: "target_player", Value: "player_2", Label: "Target Alice"},
       {Type: "target_player", Value: "player_3", Label: "Target Bob"},
   }
   ```

4. **Target Selection**: Player chooses target
   ```go
   Action{Type: "target_player", Value: "player_2"}
   ```

5. **Final Choices** (if needed): Server sends final options
   ```go
   []Action{
       {Type: "select_guess", Value: "priest", Label: "Guess Priest"},
       {Type: "select_guess", Value: "baron", Label: "Guess Baron"},
       {Type: "select_guess", Value: "handmaid", Label: "Guess Handmaid"},
       // ... etc
   }
   ```

6. **Complete Action**: Player makes final choice and action resolves

### State Updates

After each action, all players receive a `GameUpdate` containing:
- Current redacted game state
- Event describing what just happened
- Available actions for each player (redacted so players only see their own actions)

This approach allows clients to be "dumb" - they don't need to know game rules, just display state and present choices.

### Hot-Seat Support

By distinguishing public vs private state, clients can implement local multiplayer by showing/hiding private information as players take turns.

## Implementation Notes

- Use interfaces for card behaviors while maintaining data-driven state
- Separate game logic from networking/UI concerns  
- Design for extensibility (new card types, variants)
- Thread-safe for concurrent player actions
- Comprehensive error handling for invalid moves