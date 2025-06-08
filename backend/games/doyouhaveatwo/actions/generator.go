package actions

import "github.com/csmith/splendid/backend/games/doyouhaveatwo/model"

type ActionGenerator interface {
	GenerateActionsForPlayer(g *model.Game, playerID model.PlayerID) []model.Action
}

type DefaultActionGenerator struct{}

func (ag *DefaultActionGenerator) GenerateActionsForPlayer(g *model.Game, playerID model.PlayerID) []model.Action {
	player := g.GetPlayer(playerID)
	if player == nil {
		return nil
	}

	// If player has a pending action, return next steps
	if player.PendingAction.Value != nil {
		return player.PendingAction.Value.NextActions(g)
	}

	// Otherwise, generate initial actions based on game state
	return ag.generateInitialActions(g, player)
}

func (ag *DefaultActionGenerator) generateInitialActions(g *model.Game, player *model.Player) []model.Action {
	var actions []model.Action

	// For now, just generate actions for cards in hand
	for _, handCard := range player.Hand {
		if handCard.VisibleTo[player.ID] {
			switch handCard.Value {
			case model.CardGuard:
				actions = append(actions, &PlayGuardAction{
					Player: player.ID,
				})
				// Add other card types as we implement them
			}
		}
	}

	return actions
}
