package inputs

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

const InputShowdown model.InputType = "showdown"

type ShowdownInput struct {
}

func (i *ShowdownInput) Type() model.InputType {
	return InputShowdown
}

func (i *ShowdownInput) PlayerID() *model.PlayerID {
	return nil
}

func (i *ShowdownInput) Apply(g *model.Game, apply func(model.Event)) error {
	// Get all active players (not eliminated)
	activePlayers := make([]model.PlayerID, 0)
	for _, player := range g.Players {
		if !player.IsOut {
			activePlayers = append(activePlayers, player.ID)
		}
	}

	// Get all players for revealing cards to everyone
	allPlayers := make([]model.PlayerID, 0)
	for _, player := range g.Players {
		allPlayers = append(allPlayers, player.ID)
	}

	// All active players reveal their hands to ALL players
	for _, sourcePlayer := range activePlayers {
		// Create target list excluding the source player
		targetPlayers := make([]model.PlayerID, 0)
		for _, targetPlayer := range allPlayers {
			if targetPlayer != sourcePlayer {
				targetPlayers = append(targetPlayers, targetPlayer)
			}
		}

		if len(targetPlayers) > 0 {
			apply(&events.HandRevealedEvent{
				SourcePlayer:  sourcePlayer,
				TargetPlayers: targetPlayers,
			})
		}
	}

	// Find the player(s) with the highest card value
	highestValue := 0
	winners := make([]model.PlayerID, 0)

	for _, playerID := range activePlayers {
		player := g.GetPlayer(playerID)
		if player != nil && len(player.Hand) > 0 {
			cardValue := player.Hand[0].Value().Value()
			if cardValue > highestValue {
				highestValue = cardValue
				winners = []model.PlayerID{playerID}
			} else if cardValue == highestValue {
				winners = append(winners, playerID)
			}
		}
	}

	// End the round with the winners
	endRoundInput := &EndRoundInput{
		Winners: winners,
	}

	return endRoundInput.Apply(g, apply)
}
