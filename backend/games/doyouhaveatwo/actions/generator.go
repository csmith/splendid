package actions

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

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

	// During setup phase, allow adding players and starting game
	if g.Phase == model.PhaseSetup {
		if len(g.Players) < 4 {
			actions = append(actions, &AddPlayerAction{
				NewPlayerID: model.PlayerID(ag.generateRandomID()),
			})
		}
		if len(g.Players) >= 2 {
			actions = append(actions, &StartGameAction{
				Player: player.ID,
			})
		}
	}

	// During draw phase, current player can draw a card
	if g.Phase == model.PhaseDraw && g.Players[g.CurrentPlayer].ID == player.ID {
		actions = append(actions, &DrawCardAction{
			Player: player.ID,
		})
	}

	// During play phase, generate actions for cards in hand
	if g.Phase == model.PhasePlay && g.Players[g.CurrentPlayer].ID == player.ID {
		for _, handCard := range player.Hand {
			if handCard.VisibleTo[player.ID] {
				switch handCard.Value {
				case model.CardGuard:
					actions = append(actions, &PlayCardGuardAction{
						Player: player.ID,
					})
				case model.CardHandmaid, model.CardCountess, model.CardPrincess:
					actions = append(actions, &PlayCardNoTargetAction{
						Player: player.ID,
						Card:   handCard.Value,
					})
				case model.CardBaron, model.CardPriest, model.CardKing:
					actions = append(actions, &PlayCardTargetOthersAction{
						Player: player.ID,
						Card:   handCard.Value,
					})
				case model.CardPrince:
					actions = append(actions, &PlayCardTargetAnyAction{
						Player: player.ID,
						Card:   handCard.Value,
					})
				}
			}
		}
	}

	return actions
}

func (ag *DefaultActionGenerator) generateRandomID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
