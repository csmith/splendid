package model

import (
	coremodel "github.com/csmith/splendid/backend/model"
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
	Players          []*Player                    `json:"players"`
	Deck             []coremodel.Redactable[Card] `json:"deck"`
	RemovedCard      *coremodel.Redactable[Card]  `json:"removed_card"`
	SetAsideCards    []Card                       `json:"set_aside_cards"`
	CurrentPlayer    int                          `json:"current_player"`
	Round            int                          `json:"round"`
	Phase            GamePhase                    `json:"phase"`
	TokensToWin      int                          `json:"tokens_to_win"`
	LastRoundWinners []PlayerID                   `json:"last_round_winners"`
}

func (g *Game) GetPlayer(playerID PlayerID) *Player {
	for _, player := range g.Players {
		if player.ID == playerID {
			return player
		}
	}
	return nil
}
