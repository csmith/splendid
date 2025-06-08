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

- **`model/`** - Core types and data structures (Game, Player, Card, Event, Action, Input interfaces)
- **`inputs/`** - Input types containing game logic (StartRoundInput, DrawCardInput, etc.)  
- **`events/`** - Atomic event types for state mutations (DeckUpdated, CardDealt, etc.)
- **`actions/`** - Multi-step action types for complex player interactions (PlayGuardAction, etc.)
- **Main package** - Game engine and orchestration logic

### Input/Event Architecture

The game uses a sophisticated event sourcing pattern that separates business logic from state mutations:

#### Inputs
**Inputs** contain all game logic and validation. They read the current game state (immutably) and generate a sequence of atomic events to apply.

```go
type Input interface {
    Apply(g *Game, apply func(Event)) error  // Read-only Game, applies events via callback
    Type() InputType
    PlayerID() *PlayerID
}
```

**Examples:**
- `StartRoundInput` - Shuffles deck, deals cards, resets player states
- `DrawCardInput` - Validates player/deck, generates card deal event
- `PlayGuardInput` - Handles Guard card effects, targeting, eliminations

#### Events  
**Events** are atomic, immutable state mutations. They contain no logic and always succeed when applied.

```go
type Event interface {
    Type() EventType
    PlayerID() *PlayerID
    Apply(g *Game) error  // Pure state mutation, cannot fail
}
```

**Atomic Event Types:**
- `DeckUpdated` - Replace entire deck
- `CardDealt` - Move card from deck to player hand
- `CardRemoved` - Remove card to removed pile
- `PlayerEliminated` - Mark player as eliminated  
- `PlayerRestored` - Mark player as active
- `PlayerProtectionCleared` - Remove handmaid protection
- `PlayerDiscardPileCleared` - Empty discard pile
- `RoundUpdated` - Set round number
- `PhaseUpdated` - Change game phase
- `PlayerActionStarted` - Player begins multi-step action
- `PlayerActionUpdated` - Player progresses through action steps
- `PlayerActionCompleted` - Player completes action, clears pending state

#### Engine Flow
```
Action → Engine.ProcessAction() → Events → State + GameUpdate
Input → Engine.processInput() → Events → State + GameUpdate
```

1. **Action Processing**: `ProcessAction()` handles multi-step actions, storing state in `Player.PendingAction`
2. **Input Processing**: `processInput()` executes completed actions or direct inputs via callback
3. **Event Application**: Each event applies atomically via `event.Apply(game)`  
4. **History Tracking**: Events stored in `Engine.EventHistory` for replay/debugging
5. **Client Updates**: Each event triggers a `GameUpdate` to all clients

### Actions Architecture

The game implements a sophisticated actions system for handling multi-step player interactions, particularly for complex card effects that require targeting and parameter selection.

#### Actions Interface

**Actions** represent multi-step player interactions that may require several rounds of client-server communication to complete:

```go
type Action interface {
    PlayerID() PlayerID
    IsComplete() bool
    NextActions(*Game) []Action  // What choices are available next
    ToInput() Input              // Convert to concrete input when complete
    Type() string               // Action type identifier
}
```

#### Action Types

**Multi-Step Actions**: Complex cards requiring multiple choices
- `PlayGuardAction` - Target selection → Card guess → Execution
- Future: `PlayPriestAction`, `PlayBaronAction`, etc.

**Single-Step Actions**: Simple cards executed immediately
- Actions that are complete upon creation bypass pending state

#### Action State Management

**Pending Actions**: Stored in `Player.PendingAction` as `Redactable[Action]`
- Only visible to the owning player
- Tracked through events for replay/debugging
- Cleared automatically when action completes

**Action Events**:
- `PlayerActionStartedEvent` - Begin multi-step action
- `PlayerActionUpdatedEvent` - Update action with new parameters  
- `PlayerActionCompletedEvent` - Clear pending action, execute input

#### Action Flow Example (Guard Card)

1. **Initial Action**: Client calls `Engine.ProcessAction(playerID, &PlayGuardAction{Player: p1})`
2. **Start Event**: Engine generates `PlayerActionStartedEvent`, stores pending action
3. **Target Selection**: Client receives `GameUpdate` with available target actions
4. **Update Action**: Client calls `ProcessAction` with target selected
5. **Update Event**: Engine generates `PlayerActionUpdatedEvent` with target info
6. **Guess Selection**: Client receives available guess actions (all cards except Guard)
7. **Complete Action**: Client makes final guess, action becomes complete
8. **Execute**: Engine clears pending action and executes `PlayGuardInput`

#### Action Generation

**ActionGenerator**: Analyzes game state to determine available actions for each player

```go
type ActionGenerator interface {
    GenerateActionsForPlayer(g *Game, playerID PlayerID) []Action
}
```

- Checks pending actions for next steps
- Generates initial actions based on cards in hand
- Validates action availability (player turn, protection, etc.)

#### Benefits

**Multi-Step Support**: Complex card interactions handled elegantly
**Type Safety**: Each action type has typed fields instead of generic maps
**State Persistence**: Pending actions survive server restarts via event replay
**Client Simplicity**: UIs just display choices, no game logic required
**Flexible Interactions**: Easy to add new card types with complex behaviors

### Combined Architecture Benefits

**Clean Separation**: Logic lives in Inputs, state mutations in Events, interactions in Actions
**Reusable Operations**: Events like `PlayerEliminated` used across many contexts  
**Perfect Debugging**: Replay any game state from event history including partial actions
**Simplified Clients**: Clients only handle atomic event rendering and action selection
**Event Sourcing**: Complete audit trail for game actions and player interactions

#### Event Field Conventions

Events use `Result` prefix for output-only fields populated during `Apply()`:

```go
type CardDealtEvent struct {
    ToPlayer        PlayerID                    // Input parameter
    ResultCardDealt Redactable[Card]           // Output for client visibility  
}
```

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
    ID            PlayerID
    Name          string
    Hand          []Redactable[Card]
    DiscardPile   []Card                 // Player's own discard pile (visible to all)
    TokenCount    int
    IsOut         bool                   // Eliminated this round
    IsProtected   bool                   // Handmaid protection
    Position      int                    // Seating order
    PendingAction Redactable[Action]     // Multi-step action in progress
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

### Event History & Debugging

The Engine maintains a complete history of all events in `Engine.EventHistory` for:

- **Game Replay**: Reconstruct any game state by replaying events up to that point
- **Debugging**: Step through exact sequence of state changes  
- **Audit Trails**: Complete record of all game actions
- **State Verification**: Validate game state consistency

Each atomic event is immutable and contains all information needed to reproduce that state change. This enables powerful debugging capabilities where any game issue can be reproduced by replaying the exact event sequence.

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
Interface for multi-step player interactions:
```go
type Action interface {
    PlayerID() PlayerID
    IsComplete() bool
    NextActions(*Game) []Action  // What choices are available next
    ToInput() Input              // Convert to concrete input when complete
    Type() string               // Action type identifier
}
```

**Example Implementation**:
```go
type PlayGuardAction struct {
    Player       PlayerID
    TargetPlayer *PlayerID  // nil until selected
    GuessedCard  *Card      // nil until selected
}
```

#### Input Interface (model/action.go)
Represents concrete game logic and validation:
```go
type Input interface {
    Apply(g *Game, apply func(Event)) error  // Read-only access, applies events via callback
    Type() InputType
    PlayerID() *PlayerID
}
```

#### Event Interface (model/event.go)
Atomic state mutations generated by inputs:
```go
type Event interface {
    Type() EventType
    PlayerID() *PlayerID
    Apply(g *Game) error  // Pure mutation, cannot fail
}
```

**Input Implementations**: Each input contains game logic in `inputs/` package:
- `inputs/start_round.go` - StartRoundInput handles round initialization
- `inputs/draw_card.go` - DrawCardInput handles card drawing
- `inputs/play_guard.go` - PlayGuardInput handles Guard card effects

**Event Implementations**: Each atomic event in `events/` package:
- `events/deck_updated.go` - DeckUpdated replaces entire deck
- `events/card_dealt.go` - CardDealt moves card to player hand  
- `events/player_eliminated.go` - PlayerEliminated marks player as out

```go
// Example atomic events:
type DeckUpdatedEvent struct {
    NewDeck []Redactable[Card]
}
type CardDealtEvent struct {
    ToPlayer        PlayerID
    ResultCardDealt Redactable[Card]  // Populated during Apply()
}
type PlayerEliminatedEvent struct {
    Player PlayerID
}
```

**Secret State Handling**: Events expose secret information through `Result` prefixed fields populated during `Apply()`. This maintains proper visibility while providing clients with complete event information.

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

The new actions system provides a cleaner approach to multi-step interactions:

#### Client Entry Points

**Engine.ProcessAction()**: Primary method for clients to interact with the game
```go
func (e *Engine) ProcessAction(playerID PlayerID, action Action) error
```

#### Flow Examples

**Simple Action** (complete immediately):
1. Client calls `ProcessAction(playerID, &PlayHandmaidAction{Player: p1})`
2. Action is complete, executes `PlayHandmaidInput` immediately
3. Client receives `GameUpdate` with results

**Multi-Step Action** (Guard card):
1. Client calls `ProcessAction(playerID, &PlayGuardAction{Player: p1})`
2. Engine stores pending action, sends `GameUpdate` with target selection actions
3. Client calls `ProcessAction(playerID, &PlayGuardAction{Player: p1, TargetPlayer: &p2})`
4. Engine updates pending action, sends `GameUpdate` with guess selection actions  
5. Client calls `ProcessAction(playerID, &PlayGuardAction{Player: p1, TargetPlayer: &p2, GuessedCard: &CardPriest})`
6. Action complete, engine executes `PlayGuardInput`, clears pending action

#### Available Actions Generation

**Automatic**: Engine populates `GameUpdate.AvailableActions` using `ActionGenerator`
- Initial card play actions based on hand contents
- Next step actions based on pending action state
- Validates targeting, protection, game phase, etc.

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