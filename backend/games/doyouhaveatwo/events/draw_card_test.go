package events

import (
	"testing"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDrawCardEvent_Type(t *testing.T) {
	event := &DrawCardEvent{}
	assert.Equal(t, EventDrawCard, event.Type())
}

func TestDrawCardEvent_PlayerID(t *testing.T) {
	playerID := model.PlayerID("A")
	event := &DrawCardEvent{Player: playerID}
	assert.Equal(t, &playerID, event.PlayerID())
}

func TestDrawCardEvent_Apply_Success(t *testing.T) {
	playerID := model.PlayerID("A")
	player := &model.Player{
		ID:          playerID,
		Hand:        []model.Redactable[model.Card]{},
		DiscardPile: []model.Card{},
	}

	card := model.Redactable[model.Card]{
		Value:     model.CardGuard,
		VisibleTo: make(map[model.PlayerID]bool),
	}

	game := &model.Game{
		Players: []*model.Player{player},
		Deck:    []model.Redactable[model.Card]{card},
	}

	event := &DrawCardEvent{Player: playerID}

	err := event.Apply(game)

	require.NoError(t, err)
	assert.Len(t, game.Deck, 0)
	assert.Len(t, player.Hand, 1)
	assert.True(t, player.Hand[0].VisibleTo[playerID])
	assert.Equal(t, card.Value, event.CardDrawn.Value)
}

func TestDrawCardEvent_Apply_PlayerNotFound(t *testing.T) {
	game := &model.Game{
		Players: []*model.Player{},
		Deck:    []model.Redactable[model.Card]{},
	}

	event := &DrawCardEvent{Player: model.PlayerID("A")}

	err := event.Apply(game)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "player with ID A not found")
}

func TestDrawCardEvent_Apply_EmptyDeck(t *testing.T) {
	playerID := model.PlayerID("A")
	player := &model.Player{
		ID:          playerID,
		Hand:        []model.Redactable[model.Card]{},
		DiscardPile: []model.Card{},
	}

	game := &model.Game{
		Players: []*model.Player{player},
		Deck:    []model.Redactable[model.Card]{},
	}

	event := &DrawCardEvent{Player: playerID}

	err := event.Apply(game)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deck is empty, cannot draw card")
}

func TestDrawCardEvent_Apply_CardVisibility(t *testing.T) {
	playerA := model.PlayerID("A")
	playerB := model.PlayerID("B")

	players := []*model.Player{
		{ID: playerA, Hand: []model.Redactable[model.Card]{}},
		{ID: playerB, Hand: []model.Redactable[model.Card]{}},
	}

	card := model.Redactable[model.Card]{
		Value:     model.CardGuard,
		VisibleTo: make(map[model.PlayerID]bool),
	}

	game := &model.Game{
		Players: players,
		Deck:    []model.Redactable[model.Card]{card},
	}

	event := &DrawCardEvent{Player: playerA}

	err := event.Apply(game)

	require.NoError(t, err)
	drawnCard := players[0].Hand[0]
	assert.True(t, drawnCard.VisibleTo[playerA])
	assert.False(t, drawnCard.VisibleTo[playerB])
}

func TestDrawCardEvent_Apply_ExposesDrawnCard(t *testing.T) {
	playerID := model.PlayerID("A")
	player := &model.Player{
		ID:   playerID,
		Hand: []model.Redactable[model.Card]{},
	}

	expectedCard := model.Redactable[model.Card]{
		Value:     model.CardGuard,
		VisibleTo: make(map[model.PlayerID]bool),
	}

	game := &model.Game{
		Players: []*model.Player{player},
		Deck:    []model.Redactable[model.Card]{expectedCard},
	}

	event := &DrawCardEvent{Player: playerID}

	err := event.Apply(game)

	require.NoError(t, err)
	assert.Equal(t, expectedCard.Value, event.CardDrawn.Value)
	assert.True(t, event.CardDrawn.VisibleTo[playerID])
}

func TestDrawCardEvent_Apply_DrawnCardVisibilityInEvent(t *testing.T) {
	playerA := model.PlayerID("A")
	playerB := model.PlayerID("B")
	playerC := model.PlayerID("C")

	players := []*model.Player{
		{ID: playerA, Hand: []model.Redactable[model.Card]{}},
		{ID: playerB, Hand: []model.Redactable[model.Card]{}},
		{ID: playerC, Hand: []model.Redactable[model.Card]{}},
	}

	card := model.Redactable[model.Card]{
		Value:     model.CardPriest,
		VisibleTo: make(map[model.PlayerID]bool),
	}

	game := &model.Game{
		Players: players,
		Deck:    []model.Redactable[model.Card]{card},
	}

	event := &DrawCardEvent{Player: playerB}

	err := event.Apply(game)

	require.NoError(t, err)

	// Check that the drawn card in the event is visible to the drawing player
	assert.True(t, event.CardDrawn.VisibleTo[playerB])

	// Check that the drawn card in the event is not visible to other players
	assert.False(t, event.CardDrawn.VisibleTo[playerA])
	assert.False(t, event.CardDrawn.VisibleTo[playerC])

	// Verify the card value is preserved
	assert.Equal(t, "Priest", event.CardDrawn.Value.Name())
}
