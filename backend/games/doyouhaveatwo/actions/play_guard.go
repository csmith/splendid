package actions

import (
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type PlayGuardAction struct {
	Player       model.PlayerID
	TargetPlayer *model.PlayerID
	GuessedCard  *model.Card
}

func (a *PlayGuardAction) PlayerID() model.PlayerID {
	return a.Player
}

func (a *PlayGuardAction) IsComplete() bool {
	return a.TargetPlayer != nil && a.GuessedCard != nil
}

func (a *PlayGuardAction) NextActions(g *model.Game) []model.Action {
	if a.TargetPlayer == nil {
		return generateTargetActions(g, a.Player)
	}
	if a.GuessedCard == nil {
		return generateGuessActions()
	}
	return nil
}

func (a *PlayGuardAction) ToInput() model.Input {
	return &inputs.PlayGuardInput{
		Player:       a.Player,
		TargetPlayer: *a.TargetPlayer,
		GuessedRank:  a.GuessedCard.Value(),
	}
}

func (a *PlayGuardAction) Type() string {
	return "play_guard"
}

// Helper functions to generate next actions
func generateTargetActions(g *model.Game, playerID model.PlayerID) []model.Action {
	// TODO: Implement - return actions for selecting valid targets
	return nil
}

func generateGuessActions() []model.Action {
	// TODO: Implement - return actions for guessing cards (all except Guard)
	return nil
}
