package events

import (
	"testing"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartRoundEvent_Type(t *testing.T) {
	event := &StartRoundEvent{}
	assert.Equal(t, EventStartRound, event.Type())
}

func TestStartRoundEvent_PlayerID(t *testing.T) {
	event := &StartRoundEvent{}
	assert.Nil(t, event.PlayerID())
}

func TestStartRoundEvent_Apply_Success(t *testing.T) {
	players := []*model.Player{
		{
			ID:          model.PlayerID("A"),
			Hand:        []model.Redactable[model.Card]{},
			DiscardPile: []model.Card{},
			IsOut:       false,
			IsProtected: false,
		},
		{
			ID:          model.PlayerID("B"),
			Hand:        []model.Redactable[model.Card]{},
			DiscardPile: []model.Card{},
			IsOut:       false,
			IsProtected: false,
		},
	}

	game := &model.Game{
		Players: players,
		Round:   0,
		Phase:   model.PhaseSetup,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)
	assert.Equal(t, 1, game.Round)
	assert.Equal(t, model.PhasePlay, game.Phase)
	assert.NotNil(t, game.RemovedCard.Value)
	assert.Greater(t, len(game.Deck), 0)

	for _, player := range players {
		assert.Len(t, player.Hand, 1)
		assert.False(t, player.IsOut)
		assert.False(t, player.IsProtected)
		assert.Len(t, player.DiscardPile, 0)
		assert.True(t, player.Hand[0].VisibleTo[player.ID])
	}
}

func TestStartRoundEvent_Apply_IncrementsRound(t *testing.T) {
	game := &model.Game{
		Players: []*model.Player{{ID: model.PlayerID("A")}},
		Round:   5,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)
	assert.Equal(t, 6, game.Round)
}

func TestStartRoundEvent_Apply_ResetsPlayerStates(t *testing.T) {
	players := []*model.Player{
		{
			ID:          model.PlayerID("A"),
			Hand:        []model.Redactable[model.Card]{},
			DiscardPile: []model.Card{},
			IsOut:       true,
			IsProtected: true,
		},
		{
			ID:          model.PlayerID("B"),
			Hand:        []model.Redactable[model.Card]{},
			DiscardPile: []model.Card{},
			IsOut:       false,
			IsProtected: true,
		},
	}

	game := &model.Game{
		Players: players,
		Round:   0,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)
	for _, player := range players {
		assert.False(t, player.IsOut)
		assert.False(t, player.IsProtected)
		assert.Len(t, player.DiscardPile, 0)
	}
}

func TestStartRoundEvent_Apply_CardVisibility(t *testing.T) {
	playerA := model.PlayerID("A")
	playerB := model.PlayerID("B")

	players := []*model.Player{
		{ID: playerA, Hand: []model.Redactable[model.Card]{}},
		{ID: playerB, Hand: []model.Redactable[model.Card]{}},
	}

	game := &model.Game{
		Players: players,
		Round:   0,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)

	playerACard := players[0].Hand[0]
	playerBCard := players[1].Hand[0]

	assert.True(t, playerACard.VisibleTo[playerA])
	assert.False(t, playerACard.VisibleTo[playerB])
	assert.True(t, playerBCard.VisibleTo[playerB])
	assert.False(t, playerBCard.VisibleTo[playerA])
}

func TestStartRoundEvent_Apply_DeckContainsExpectedCards(t *testing.T) {
	game := &model.Game{
		Players: []*model.Player{{ID: model.PlayerID("A")}},
		Round:   0,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)

	totalCards := len(game.Deck) + 1 + len(game.Players)

	assert.Equal(t, 16, totalCards)
}

func TestStartRoundEvent_Apply_RemovedCardNotVisibleToPlayers(t *testing.T) {
	players := []*model.Player{
		{ID: model.PlayerID("A")},
		{ID: model.PlayerID("B")},
	}

	game := &model.Game{
		Players: players,
		Round:   0,
	}

	event := &StartRoundEvent{}

	err := event.Apply(game)

	require.NoError(t, err)

	for playerID, visible := range game.RemovedCard.VisibleTo {
		assert.False(t, visible, "removed card should not be visible to player %s", playerID)
	}
}

func TestStartRoundEvent_createShuffledDeck(t *testing.T) {
	event := &StartRoundEvent{}
	deck := event.createShuffledDeck()

	assert.Len(t, deck, 16)

	cardCounts := make(map[string]int)
	for _, card := range deck {
		cardCounts[card.Value.Name()]++
	}

	assert.Equal(t, 5, cardCounts["Guard"])
	assert.Equal(t, 2, cardCounts["Priest"])
	assert.Equal(t, 2, cardCounts["Baron"])
	assert.Equal(t, 2, cardCounts["Handmaid"])
	assert.Equal(t, 2, cardCounts["Prince"])
	assert.Equal(t, 1, cardCounts["King"])
	assert.Equal(t, 1, cardCounts["Countess"])
	assert.Equal(t, 1, cardCounts["Princess"])

	for _, card := range deck {
		assert.NotNil(t, card.VisibleTo)
		assert.Len(t, card.VisibleTo, 0)
	}
}
