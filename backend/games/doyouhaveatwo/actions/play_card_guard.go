package actions

import (
	"fmt"
	"strings"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/csmith/splendid/backend/serialization"
)

type PlayCardWithGuessAction struct {
	Player       model.PlayerID  `json:"player"`
	CardName     string          `json:"card_name"`
	TargetPlayer *model.PlayerID `json:"target_player"`
	GuessedRank  *int            `json:"guessed_rank"`
}

func (a *PlayCardWithGuessAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardWithGuessAction) IsComplete() bool {
	return a.TargetPlayer != nil && a.GuessedRank != nil
}

func (a *PlayCardWithGuessAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	if a.GuessedRank == nil {
		return a.generateGuessActions()
	}
	return nil
}

func (a *PlayCardWithGuessAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}

	switch a.CardName {
	case "Guard":
		return &inputs.PlayGuardInput{
			Player:       a.Player,
			TargetPlayer: *a.TargetPlayer,
			GuessedRank:  *a.GuessedRank,
		}
	default:
		return nil
	}
}

func (a *PlayCardWithGuessAction) Type() serialization.Specifier {
	return specifier("play_card_with_guess")
}

func (a *PlayCardWithGuessAction) String() string {
	cardNameLower := strings.ToLower(a.CardName)
	if a.TargetPlayer == nil && a.GuessedRank == nil {
		return fmt.Sprintf("play_%s(player=%s)", cardNameLower, a.Player)
	} else if a.GuessedRank == nil {
		return fmt.Sprintf("play_%s(player=%s, target=%s)", cardNameLower, a.Player, *a.TargetPlayer)
	} else {
		return fmt.Sprintf("play_%s(player=%s, target=%s, guess=%d)", cardNameLower, a.Player, *a.TargetPlayer, *a.GuessedRank)
	}
}

func (a *PlayCardWithGuessAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayCardWithGuessAction{
				Player:       a.Player,
				CardName:     a.CardName,
				TargetPlayer: &player.ID,
				GuessedRank:  a.GuessedRank,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}

func (a *PlayCardWithGuessAction) generateGuessActions() []model.Action {
	var actions []model.Action

	// Generate actions for each valid guess (all card ranks except Guard)
	validRanks := []int{2, 3, 4, 5, 6, 7, 8} // Priest, Baron, Handmaid, Prince, King, Countess, Princess

	for _, rank := range validRanks {
		// Create a new action with this guess selected
		guessAction := &PlayCardWithGuessAction{
			Player:       a.Player,
			CardName:     a.CardName,
			TargetPlayer: a.TargetPlayer,
			GuessedRank:  &rank,
		}
		actions = append(actions, guessAction)
	}

	return actions
}
