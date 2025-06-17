package actions

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type Generator struct{}

func (ag *Generator) GenerateActionsForPlayer(g *model.Game, playerID model.PlayerID) []model.GameAction {
	player := g.GetPlayer(playerID)
	if player == nil {
		return nil
	}

	return ag.generateInitialActions(g, player)
}

func (ag *Generator) generateInitialActions(g *model.Game, player *model.Player) []model.GameAction {
	var actions []model.GameAction

	// During setup phase, allow starting game
	if g.Phase == model.PhaseSetup {
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
		// Check for Countess rule: if player has Countess and (Prince or King), must play Countess
		hasCountess := false
		hasPrinceOrKing := false
		for _, handCard := range player.Hand {
			if handCard.Value() == model.CardCountess {
				hasCountess = true
			}
			if handCard.Value() == model.CardPrince || handCard.Value() == model.CardKing {
				hasPrinceOrKing = true
			}
		}

		// If Countess rule applies, only allow Countess to be played
		if hasCountess && hasPrinceOrKing {
			actions = append(actions, &PlayCardNoTargetAction{
				Player:   player.ID,
				CardName: "Countess",
			})
		} else {
			for _, handCard := range player.Hand {
				// Note: All cards in hand are visible to the player by default
				switch handCard.Value() {
				case model.CardGuard:
					// Check if there are valid targets for Guard
					if ag.hasValidTargetsForOthers(g, player.ID) {
						actions = append(actions, &PlayCardWithGuessAction{
							Player:   player.ID,
							CardName: "Guard",
						})
					} else {
						// No valid targets, offer discard instead
						actions = append(actions, &DiscardCardAction{
							Player:   player.ID,
							CardName: handCard.Value().Name(),
						})
					}
				case model.CardHandmaid, model.CardCountess, model.CardPrincess:
					actions = append(actions, &PlayCardNoTargetAction{
						Player:   player.ID,
						CardName: handCard.Value().Name(),
					})
				case model.CardBaron, model.CardPriest, model.CardKing:
					// Check if there are valid targets for these cards
					if ag.hasValidTargetsForOthers(g, player.ID) {
						actions = append(actions, &PlayCardTargetOthersAction{
							Player:   player.ID,
							CardName: handCard.Value().Name(),
						})
					} else {
						// No valid targets, offer discard instead
						actions = append(actions, &DiscardCardAction{
							Player:   player.ID,
							CardName: handCard.Value().Name(),
						})
					}
				case model.CardPrince:
					// Prince can always target self if no others available
					actions = append(actions, &PlayCardTargetAnyAction{
						Player:   player.ID,
						CardName: handCard.Value().Name(),
					})
				}
			}
		}
	}

	return actions
}

// hasValidTargetsForOthers checks if there are any valid targets for cards that target other players
func (ag *Generator) hasValidTargetsForOthers(g *model.Game, playerID model.PlayerID) bool {
	for _, player := range g.Players {
		// Can target other players who are not eliminated and not protected
		if player.ID != playerID && !player.IsOut && !player.IsProtected {
			return true
		}
	}
	return false
}
