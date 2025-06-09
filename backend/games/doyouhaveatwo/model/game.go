package model

import (
	"github.com/csmith/splendid/backend/serialization"
)

type GamePhase string

const (
	PhaseSetup    GamePhase = "setup"
	PhaseDraw     GamePhase = "draw"
	PhasePlay     GamePhase = "play"
	PhaseRoundEnd GamePhase = "round_end"
	PhaseGameEnd  GamePhase = "game_end"
)

type Game struct {
	Players          []*Player                        `json:"players"`
	Deck             []serialization.Redactable[Card] `json:"deck"`
	RemovedCard      *serialization.Redactable[Card]  `json:"removed_card"`
	CurrentPlayer    int                              `json:"current_player"`
	Round            int                              `json:"round"`
	Phase            GamePhase                        `json:"phase"`
	TokensToWin      int                              `json:"tokens_to_win"`
	LastRoundWinners []PlayerID                       `json:"last_round_winners"`
}

func (g *Game) GetPlayer(playerID PlayerID) *Player {
	for _, player := range g.Players {
		if player.ID == playerID {
			return player
		}
	}
	return nil
}
