package actions

import (
	"encoding/json"
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayCardGuardAction struct {
	Player       model.PlayerID  `json:"player"`
	TargetPlayer *model.PlayerID `json:"target_player"`
	GuessedRank  *int            `json:"guessed_rank"`
}

func (a *PlayCardGuardAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayCardGuardAction) IsComplete() bool {
	return a.TargetPlayer != nil && a.GuessedRank != nil
}

func (a *PlayCardGuardAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return a.generateTargetActions(g)
	}
	if a.GuessedRank == nil {
		return a.generateGuessActions()
	}
	return nil
}

func (a *PlayCardGuardAction) ToInput() model.Input {
	if !a.IsComplete() {
		return nil
	}
	return &inputs.PlayGuardInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
		GuessedRank:  *a.GuessedRank,
	}
}

func (a *PlayCardGuardAction) Type() string {
	return "play_guard"
}

func (a *PlayCardGuardAction) String() string {
	if a.TargetPlayer == nil && a.GuessedRank == nil {
		return fmt.Sprintf("play_guard(player=%s)", a.Player)
	} else if a.GuessedRank == nil {
		return fmt.Sprintf("play_guard(player=%s, target=%s)", a.Player, *a.TargetPlayer)
	} else {
		return fmt.Sprintf("play_guard(player=%s, target=%s, guess=%d)", a.Player, *a.TargetPlayer, *a.GuessedRank)
	}
}

func (a *PlayCardGuardAction) generateTargetActions(g *model.Game) []model.Action {
	var actions []model.Action

	// Generate actions for each valid target player
	for _, player := range g.Players {
		// Can't target self, eliminated players, or protected players
		if player.ID != a.Player && !player.IsOut && !player.IsProtected {
			// Create a new action with this target selected
			targetAction := &PlayCardGuardAction{
				Player:       a.Player,
				TargetPlayer: &player.ID,
				GuessedRank:  a.GuessedRank,
			}
			actions = append(actions, targetAction)
		}
	}

	return actions
}

func (a *PlayCardGuardAction) generateGuessActions() []model.Action {
	var actions []model.Action

	// Generate actions for each valid guess (all card ranks except Guard)
	validRanks := []int{2, 3, 4, 5, 6, 7, 8} // Priest, Baron, Handmaid, Prince, King, Countess, Princess

	for _, rank := range validRanks {
		// Create a new action with this guess selected
		guessAction := &PlayCardGuardAction{
			Player:       a.Player,
			TargetPlayer: a.TargetPlayer,
			GuessedRank:  &rank,
		}
		actions = append(actions, guessAction)
	}

	return actions
}

func (a *PlayCardGuardAction) MarshalJSON() ([]byte, error) {
	type Alias PlayCardGuardAction
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  a.Type(),
		Alias: (*Alias)(a),
	})
}
