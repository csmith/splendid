package model

type GamePhase string

const (
	PhaseSetup    GamePhase = "setup"
	PhasePlay     GamePhase = "play"
	PhaseRoundEnd GamePhase = "round_end"
	PhaseGameEnd  GamePhase = "game_end"
)

type Game struct {
	Players       []*Player
	Deck          []Redactable[Card]
	RemovedCard   Redactable[Card]
	CurrentPlayer int
	Round         int
	Phase         GamePhase
	TokensToWin   int
}

func (g *Game) GetPlayer(playerID PlayerID) *Player {
	for _, player := range g.Players {
		if player.ID == playerID {
			return player
		}
	}
	return nil
}