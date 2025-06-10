package inputs

import (
	"fmt"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

type AddPlayerInput struct {
	NewPlayerID   model.PlayerID
	NewPlayerName string
}

func (i *AddPlayerInput) PlayerID() *model.PlayerID {
	return nil
}

func (i *AddPlayerInput) Apply(g *model.Game, apply func(model.Event)) error {
	if g.GetPlayer(i.NewPlayerID) != nil {
		return fmt.Errorf("player with ID %s already exists", i.NewPlayerID)
	}

	if g.Phase != model.PhaseSetup {
		return fmt.Errorf("can only add players during setup phase")
	}

	if len(g.Players) >= 4 {
		return fmt.Errorf("maximum number of players (4) already reached")
	}

	if i.NewPlayerName == "" {
		return fmt.Errorf("player name cannot be empty")
	}

	apply(&events.PlayerAddedEvent{
		NewPlayerID:   i.NewPlayerID,
		NewPlayerName: i.NewPlayerName,
	})

	return nil
}
