package events

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	coremodel "github.com/csmith/splendid/backend/model"
)

type LastRoundWinnersUpdatedEvent struct {
	Winners []model.PlayerID `json:"winners"`
}

func (e *LastRoundWinnersUpdatedEvent) Type() coremodel.Specifier {
	return specifier("last_round_winners_updated")
}

func (e *LastRoundWinnersUpdatedEvent) PlayerID() *model.PlayerID {
	return nil
}

func (e *LastRoundWinnersUpdatedEvent) Apply(g *model.Game) error {
	g.LastRoundWinners = e.Winners
	return nil
}
