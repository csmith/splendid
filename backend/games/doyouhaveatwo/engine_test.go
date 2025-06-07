package doyouhaveatwo

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

type EngineTestSuite struct {
	engine      *Engine
	game        *Game
	players     []*Player
	updateChan  chan GameUpdate
	eventChan   chan Event
	lastUpdate  GameUpdate
	initialDeck []Redactable[Card]
	lastError   error
}

func (s *EngineTestSuite) givenAGameWithPlayers(playerCount int) error {
	s.players = make([]*Player, playerCount)
	for i := 0; i < playerCount; i++ {
		s.players[i] = &Player{
			ID:          PlayerID(rune('A' + i)),
			Name:        string(rune('A' + i)),
			Hand:        []Redactable[Card]{},
			DiscardPile: []Card{},
			TokenCount:  0,
			IsOut:       false,
			IsProtected: false,
			Position:    i,
		}
	}

	s.game = &Game{
		Players:       s.players,
		Deck:          []Redactable[Card]{},
		CurrentPlayer: 0,
		Round:         0,
		Phase:         PhaseSetup,
		TokensToWin:   5,
	}

	s.updateChan = make(chan GameUpdate, 100)
	s.eventChan = make(chan Event, 100)

	s.engine = &Engine{
		Game:       *s.game,
		updateChan: s.updateChan,
		eventChan:  s.eventChan,
	}

	return nil
}

func (s *EngineTestSuite) givenTheGameHasPhase(phaseName string) error {
	s.engine.Game.Phase = GamePhase(phaseName)
	return nil
}

func (s *EngineTestSuite) givenTheCurrentRoundIs(roundNumber int) error {
	s.engine.Game.Round = roundNumber
	return nil
}

func (s *EngineTestSuite) whenARoundStarts() error {
	s.initialDeck = make([]Redactable[Card], len(s.engine.Game.Deck))
	copy(s.initialDeck, s.engine.Game.Deck)

	event := Event{Type: EventStartRound}
	s.engine.applyEvent(event)

	// Capture the game update
	select {
	case update := <-s.updateChan:
		s.lastUpdate = update
	default:
		// No update received
	}

	return nil
}

func (s *EngineTestSuite) thenTheRoundNumberShouldBe(expectedRound int) error {
	if s.engine.Game.Round != expectedRound {
		return fmt.Errorf("expected round number to be %d, but got %d", expectedRound, s.engine.Game.Round)
	}
	return nil
}

func (s *EngineTestSuite) thenThereAreCardCardsInTheGame(expectedCount int, cardType string) error {
	var allCards []Card

	// Collect all cards from deck
	for _, card := range s.engine.Game.Deck {
		if card.Value != nil {
			allCards = append(allCards, card.Value)
		}
	}

	// Collect removed card
	if s.engine.Game.RemovedCard.Value != nil {
		allCards = append(allCards, s.engine.Game.RemovedCard.Value)
	}

	// Collect all cards from player hands
	for _, player := range s.engine.Game.Players {
		for _, card := range player.Hand {
			if card.Value != nil {
				allCards = append(allCards, card.Value)
			}
		}
	}

	// Collect all cards from player discard piles
	for _, player := range s.engine.Game.Players {
		allCards = append(allCards, player.DiscardPile...)
	}

	// Count cards of the specified type
	count := 0
	for _, card := range allCards {
		if card.Name() == cardType {
			count++
		}
	}

	if count != expectedCount {
		return fmt.Errorf("expected %d %s cards in the game, but got %d", expectedCount, cardType, count)
	}
	return nil
}

func (s *EngineTestSuite) thenACardShouldBeRemovedFromTheGame() error {
	if s.engine.Game.RemovedCard.Value == nil {
		return fmt.Errorf("expected a card to be removed from the game, but no card was removed")
	}
	return nil
}

func (s *EngineTestSuite) thenTheRemovedCardShouldNotBeVisibleToAnyPlayer() error {
	for playerID := range s.engine.Game.RemovedCard.VisibleTo {
		if s.engine.Game.RemovedCard.VisibleTo[playerID] {
			return fmt.Errorf("expected removed card to not be visible to any player, but it is visible to player %s", playerID)
		}
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldHaveExactlyCardInTheirHand(playerID string, cardCount int) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if len(player.Hand) != cardCount {
		return fmt.Errorf("expected player %s to have %d cards in hand, but got %d", player.ID, cardCount, len(player.Hand))
	}
	return nil
}

func (s *EngineTestSuite) thenPlayersCardsShouldOnlyBeVisibleToThemselves(playerID string) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	for _, card := range player.Hand {
		// Check that the card is visible to the player
		if !card.VisibleTo[player.ID] {
			return fmt.Errorf("expected player %s's card to be visible to themselves, but it is not", player.ID)
		}
		// Check that the card is not visible to other players
		for otherPlayerID, visible := range card.VisibleTo {
			if otherPlayerID != player.ID && visible {
				return fmt.Errorf("expected player %s's card to only be visible to themselves, but it is visible to player %s", player.ID, otherPlayerID)
			}
		}
	}
	return nil
}

func (s *EngineTestSuite) givenPlayerHasCardsInDiscardPile(playerID string, cardCount int) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	cards := make([]Card, cardCount)
	for i := 0; i < cardCount; i++ {
		cards[i] = Guard{} // Use Guard as placeholder
	}
	player.DiscardPile = cards
	return nil
}

func (s *EngineTestSuite) givenPlayerIsEliminated(playerID string) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	player.IsOut = true
	return nil
}

func (s *EngineTestSuite) givenPlayerIsProtected(playerID string) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	player.IsProtected = true
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldNotBeEliminated(playerID string) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if player.IsOut {
		return fmt.Errorf("expected player %s to not be eliminated, but they are", player.ID)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldNotBeProtected(playerID string) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if player.IsProtected {
		return fmt.Errorf("expected player %s to not be protected, but they are", player.ID)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldHaveCardsInDiscardPile(playerID string, expectedCount int) error {
	player := s.engine.Game.GetPlayer(PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if len(player.DiscardPile) != expectedCount {
		return fmt.Errorf("expected player %s to have %d cards in discard pile, but got %d", player.ID, expectedCount, len(player.DiscardPile))
	}
	return nil
}

func (s *EngineTestSuite) thenTheGamePhaseShouldBe(phaseName string) error {
	if string(s.engine.Game.Phase) != phaseName {
		return fmt.Errorf("expected game phase to be %v, but got %v", phaseName, s.engine.Game.Phase)
	}
	return nil
}

func (s *EngineTestSuite) thenTheRoundNumberShouldBeIncremented() error {
	if s.engine.Game.Round == 0 {
		return fmt.Errorf("expected round number to be incremented from 0, but it is still 0")
	}
	return nil
}

func (s *EngineTestSuite) thenTheDeckShouldBeCreatedAndShuffled() error {
	if len(s.engine.Game.Deck) == 0 {
		return fmt.Errorf("expected deck to be created and have cards, but deck is empty")
	}
	return nil
}

func (s *EngineTestSuite) thenTheTopCardShouldBeRemovedFromDeck() error {
	if s.engine.Game.RemovedCard.Value == nil {
		return fmt.Errorf("expected a card to be removed from deck, but removed card is nil")
	}
	return nil
}

func (s *EngineTestSuite) thenEachPlayerShouldReceiveOneCard() error {
	for _, player := range s.engine.Game.Players {
		if len(player.Hand) != 1 {
			return fmt.Errorf("expected player %s to receive one card, but got %d cards", player.ID, len(player.Hand))
		}
	}
	return nil
}

func (s *EngineTestSuite) thenAllPlayerStatesShouldBeReset() error {
	for _, player := range s.engine.Game.Players {
		if player.IsOut || player.IsProtected || len(player.DiscardPile) != 0 {
			return fmt.Errorf("expected player %s states to be reset, but IsOut=%v, IsProtected=%v, DiscardPile=%d", player.ID, player.IsOut, player.IsProtected, len(player.DiscardPile))
		}
	}
	return nil
}

func (s *EngineTestSuite) thenTheDeckShouldHaveCardsRemaining(expectedCount int) error {
	actualCount := len(s.engine.Game.Deck)
	if actualCount != expectedCount {
		return fmt.Errorf("expected deck to have %d cards remaining, but got %d", expectedCount, actualCount)
	}
	return nil
}

func (s *EngineTestSuite) whenPlayerDrawsACard(playerID string) error {
	event := Event{
		Type:     EventDrawCard,
		PlayerID: (*PlayerID)(&playerID),
	}
	s.lastError = s.engine.applyEvent(event)
	return nil // Don't propagate the error to godog, we'll check it in the Then step
}

func (s *EngineTestSuite) givenTheDeckIsEmpty() error {
	s.engine.Game.Deck = []Redactable[Card]{}
	return nil
}

func (s *EngineTestSuite) thenAnErrorIsReturned(expectedError string) error {
	if s.lastError == nil {
		return fmt.Errorf("expected error '%s' but no error was returned", expectedError)
	}
	if s.lastError.Error() != expectedError {
		return fmt.Errorf("expected error '%s' but got '%s'", expectedError, s.lastError.Error())
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	suite := &EngineTestSuite{}

	ctx.Given(`^a game with (\d+) players$`, suite.givenAGameWithPlayers)
	ctx.Given(`^the game has phase "([^"]*)"$`, suite.givenTheGameHasPhase)
	ctx.Given(`^the current round is (\d+)$`, suite.givenTheCurrentRoundIs)
	ctx.Given(`^a round starts$`, suite.whenARoundStarts)
	ctx.Given(`^player ([A-Z]) has (\d+) cards in discard pile$`, suite.givenPlayerHasCardsInDiscardPile)
	ctx.Given(`^player ([A-Z]) is eliminated$`, suite.givenPlayerIsEliminated)
	ctx.Given(`^player ([A-Z]) is protected$`, suite.givenPlayerIsProtected)
	ctx.Given(`^the deck is empty$`, suite.givenTheDeckIsEmpty)

	ctx.When(`^a round starts$`, suite.whenARoundStarts)
	ctx.When(`^player ([A-Z]) draws a card$`, suite.whenPlayerDrawsACard)

	ctx.Then(`^the round number should be (\d+)$`, suite.thenTheRoundNumberShouldBe)
	ctx.Then(`^there are (\d+) ([A-Za-z]+) cards in the game$`, suite.thenThereAreCardCardsInTheGame)
	ctx.Then(`^a card should be removed from the game$`, suite.thenACardShouldBeRemovedFromTheGame)
	ctx.Then(`^the removed card should not be visible to any player$`, suite.thenTheRemovedCardShouldNotBeVisibleToAnyPlayer)
	ctx.Then(`^player ([A-Z]) should have exactly (\d+) card in their hand$`, suite.thenPlayerShouldHaveExactlyCardInTheirHand)
	ctx.Then(`^player ([A-Z])'s cards should only be visible to themselves$`, suite.thenPlayersCardsShouldOnlyBeVisibleToThemselves)
	ctx.Then(`^player ([A-Z]) should not be eliminated$`, suite.thenPlayerShouldNotBeEliminated)
	ctx.Then(`^player ([A-Z]) should not be protected$`, suite.thenPlayerShouldNotBeProtected)
	ctx.Then(`^player ([A-Z]) should have (\d+) cards in discard pile$`, suite.thenPlayerShouldHaveCardsInDiscardPile)
	ctx.Then(`^the game phase should be "([^"]*)"$`, suite.thenTheGamePhaseShouldBe)
	ctx.Then(`^the round number should be incremented$`, suite.thenTheRoundNumberShouldBeIncremented)
	ctx.Then(`^the deck should be created and shuffled$`, suite.thenTheDeckShouldBeCreatedAndShuffled)
	ctx.Then(`^the top card should be removed from deck$`, suite.thenTheTopCardShouldBeRemovedFromDeck)
	ctx.Then(`^each player should receive one card$`, suite.thenEachPlayerShouldReceiveOneCard)
	ctx.Then(`^all player states should be reset$`, suite.thenAllPlayerStatesShouldBeReset)
	ctx.Then(`^the deck should have (\d+) cards remaining$`, suite.thenTheDeckShouldHaveCardsRemaining)
	ctx.Then(`^an error is returned: "([^"]*)"$`, suite.thenAnErrorIsReturned)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Strict:   true,
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
