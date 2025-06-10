package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type EndTurnInput struct {
	Player model.PlayerID
}

func (i *EndTurnInput) PlayerID() *model.PlayerID {
	return &i.Player
}

func (i *EndTurnInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Count remaining active players (now sees the current state with eliminations applied!)
	activePlayers := []*model.Player{}
	for _, player := range g.Players {
		if !player.IsOut {
			activePlayers = append(activePlayers, player)
		}
	}

	// If only one player remains, they win the round
	if len(activePlayers) == 1 {
		// End the round
		endRoundInput := &EndRoundInput{
			Winners: []model.PlayerID{activePlayers[0].ID},
		}

		return endRoundInput.Apply(g, apply)
	} else {
		// Round continues - advance to next player
		// Find next active player
		nextPlayerIndex := (g.CurrentPlayer + 1) % len(g.Players)
		for g.Players[nextPlayerIndex].IsOut {
			nextPlayerIndex = (nextPlayerIndex + 1) % len(g.Players)
		}

		// Update current player
		apply(&events.CurrentPlayerUpdatedEvent{
			NewCurrentPlayer: nextPlayerIndex,
		})

		// Set phase to draw for the next player's turn
		apply(&events.PhaseUpdatedEvent{
			NewPhase: model.PhaseDraw,
		})

		// If the new player is protected, their protection wanes
		if g.Players[nextPlayerIndex].IsProtected {
			apply(&events.PlayerProtectionClearedEvent{
				Player: g.Players[nextPlayerIndex].ID,
			})
		}
	}

	return nil
}
