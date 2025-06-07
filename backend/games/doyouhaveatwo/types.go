package doyouhaveatwo

type PlayerID string

type GamePhase string

const (
	PhaseSetup    GamePhase = "setup"
	PhasePlay     GamePhase = "play"
	PhaseRoundEnd GamePhase = "round_end"
	PhaseGameEnd  GamePhase = "game_end"
)

type Redactable[T any] struct {
	Value     T
	VisibleTo map[PlayerID]bool
}

type Player struct {
	ID          PlayerID
	Name        string
	Hand        []Redactable[Card]
	DiscardPile []Card
	TokenCount  int
	IsOut       bool
	IsProtected bool
	Position    int
}

type Game struct {
	Players       []*Player
	Deck          []Redactable[Card]
	RemovedCard   Redactable[Card]
	CurrentPlayer int
	Round         int
	Phase         GamePhase
	TokensToWin   int
}

type Action struct {
	Type  string
	Value interface{}
	Label string
}

type GameUpdate struct {
	Game             Game
	Event            Event
	AvailableActions map[PlayerID]Redactable[[]Action]
}
