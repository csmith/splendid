package doyouhaveatwo

import (
	"fmt"
	"testing"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/events"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/actions"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/inputs"
	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
	"github.com/cucumber/godog"
)

type EngineTestSuite struct {
	engine      *Engine
	game        *model.Game
	players     []*model.Player
	updateChan  chan model.GameUpdate
	eventChan   chan model.Event
	lastUpdate  model.GameUpdate
	initialDeck []model.Redactable[model.Card]
	lastError   error
}

func (s *EngineTestSuite) dumpEvents() string {
	if s.engine == nil || len(s.engine.EventHistory) == 0 {
		return "No events recorded"
	}

	result := "Event History:\n"
	for i, event := range s.engine.EventHistory {
		result += fmt.Sprintf("  %d: %s - %+v\n", i+1, event.Type(), event)
	}
	return result
}

func (s *EngineTestSuite) wrapError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%v\n%s", err, s.dumpEvents())
}

func (s *EngineTestSuite) givenAGameWithPlayers(playerCount int) error {
	s.players = make([]*model.Player, playerCount)
	for i := 0; i < playerCount; i++ {
		s.players[i] = &model.Player{
			ID:          model.PlayerID(rune('A' + i)),
			Name:        string(rune('A' + i)),
			Hand:        []model.Redactable[model.Card]{},
			DiscardPile: []model.Card{},
			TokenCount:  0,
			IsOut:       false,
			IsProtected: false,
			Position:    i,
		}
	}

	s.game = &model.Game{
		Players:       s.players,
		Deck:          []model.Redactable[model.Card]{},
		CurrentPlayer: 0,
		Round:         0,
		Phase:         model.PhaseSetup,
		TokensToWin:   5,
	}

	s.updateChan = make(chan model.GameUpdate, 100)
	s.eventChan = make(chan model.Event, 100)

	s.engine = &Engine{
		Game:            *s.game,
		EventHistory:    []model.Event{},
		updateChan:      s.updateChan,
		actionGenerator: &actions.DefaultActionGenerator{},
	}

	return nil
}

func (s *EngineTestSuite) givenTheGameHasPhase(phaseName string) error {
	s.engine.Game.Phase = model.GamePhase(phaseName)
	return nil
}

func (s *EngineTestSuite) givenTheCurrentRoundIs(roundNumber int) error {
	s.engine.Game.Round = roundNumber
	return nil
}

func (s *EngineTestSuite) whenARoundStarts() error {
	s.initialDeck = make([]model.Redactable[model.Card], len(s.engine.Game.Deck))
	copy(s.initialDeck, s.engine.Game.Deck)

	input := &inputs.StartRoundInput{}
	s.engine.processInput(input)

	// Capture the last game update
	var lastUpdate model.GameUpdate
	updateReceived := false
	for {
		select {
		case update := <-s.updateChan:
			lastUpdate = update
			updateReceived = true
		default:
			goto done
		}
	}
done:
	if updateReceived {
		s.lastUpdate = lastUpdate
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
	var allCards []model.Card

	// Collect all cards from deck
	for _, card := range s.engine.Game.Deck {
		allCards = append(allCards, card.Value)
	}

	// Collect removed card
	if s.engine.Game.RemovedCard != nil {
		allCards = append(allCards, s.engine.Game.RemovedCard.Value)
	}

	// Collect all cards from player hands
	for _, player := range s.engine.Game.Players {
		for _, card := range player.Hand {
			allCards = append(allCards, card.Value)
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
	if s.engine.Game.RemovedCard == nil {
		return fmt.Errorf("expected a card to be removed from the game, but no card was removed")
	}
	return nil
}

func (s *EngineTestSuite) thenTheRemovedCardShouldNotBeVisibleToAnyPlayer() error {
	if s.engine.Game.RemovedCard == nil {
		return fmt.Errorf("no card was removed")
	}
	for playerID := range s.engine.Game.RemovedCard.VisibleTo {
		if s.engine.Game.RemovedCard.VisibleTo[playerID] {
			return fmt.Errorf("expected removed card to not be visible to any player, but it is visible to player %s", playerID)
		}
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldHaveExactlyCardInTheirHand(playerID string, cardCount int) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if len(player.Hand) != cardCount {
		return fmt.Errorf("expected player %s to have %d cards in hand, but got %d", player.ID, cardCount, len(player.Hand))
	}
	return nil
}

func (s *EngineTestSuite) thenPlayersCardsShouldOnlyBeVisibleToThemselves(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
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
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	discards := make([]model.Card, cardCount)
	for i := 0; i < cardCount; i++ {
		discards[i] = model.CardGuard // Use Guard as placeholder
	}
	player.DiscardPile = discards
	return nil
}

func (s *EngineTestSuite) givenPlayerIsEliminated(playerID string) error {
	return s.engine.applyEvent(&events.PlayerEliminatedEvent{Player: model.PlayerID(playerID)})
}

func (s *EngineTestSuite) givenPlayerIsProtected(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	player.IsProtected = true
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldNotBeEliminated(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if player.IsOut {
		return fmt.Errorf("expected player %s to not be eliminated, but they are", player.ID)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldNotBeProtected(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if player.IsProtected {
		return fmt.Errorf("expected player %s to not be protected, but they are", player.ID)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldHaveCardsInDiscardPile(playerID string, expectedCount int) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return s.wrapError(fmt.Errorf("player %s not found", playerID))
	}
	if len(player.DiscardPile) != expectedCount {
		return s.wrapError(fmt.Errorf("expected player %s to have %d cards in discard pile, but got %d", player.ID, expectedCount, len(player.DiscardPile)))
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
	if s.engine.Game.RemovedCard == nil {
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
	input := &inputs.DrawCardInput{
		Player: model.PlayerID(playerID),
	}
	s.lastError = s.engine.processInput(input)
	return nil // Don't propagate the error to godog, we'll check it in the Then step
}

func (s *EngineTestSuite) givenTheDeckIsEmpty() error {
	s.engine.Game.Deck = []model.Redactable[model.Card]{}
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

func (s *EngineTestSuite) thenAnErrorOccurs() error {
	if s.lastError == nil {
		return s.wrapError(fmt.Errorf("expected an error to occur but no error was returned"))
	}
	return nil
}

func (s *EngineTestSuite) thenNoErrorOccurs() error {
	if s.lastError != nil {
		return s.wrapError(fmt.Errorf("expected no error but got: %v", s.lastError))
	}
	return nil
}

func (s *EngineTestSuite) whenPlayerPerformsActionAddPlayer(playerID, newPlayerID, newPlayerName string) error {
	action := &actions.AddPlayerAction{
		NewPlayerID:   model.PlayerID(newPlayerID),
		NewPlayerName: newPlayerName,
	}
	s.lastError = s.engine.ProcessAction(model.PlayerID(playerID), action)
	return nil
}

func (s *EngineTestSuite) whenPlayerPerformsActionStartGame(playerID string) error {
	action := &actions.StartGameAction{
		Player: model.PlayerID(playerID),
	}
	s.lastError = s.engine.ProcessAction(model.PlayerID(playerID), action)
	return nil
}

func (s *EngineTestSuite) whenPlayerPerformsActionPlayGuardTargetingGuessing(playerID, targetPlayerID string, guessedRank int) error {
	playerIDTyped := model.PlayerID(playerID)
	targetPlayerIDPtr := model.PlayerID(targetPlayerID)

	// Step 1: Start the play_guard action
	initialAction := &actions.PlayGuardAction{
		Player: playerIDTyped,
	}
	s.lastError = s.engine.ProcessAction(playerIDTyped, initialAction)
	if s.lastError != nil {
		return nil
	}

	// Step 2: Select target player
	targetAction := &actions.PlayGuardAction{
		Player:       playerIDTyped,
		TargetPlayer: &targetPlayerIDPtr,
	}
	s.lastError = s.engine.ProcessAction(playerIDTyped, targetAction)
	if s.lastError != nil {
		return nil
	}

	// Step 3: Select guessed rank
	finalAction := &actions.PlayGuardAction{
		Player:       playerIDTyped,
		TargetPlayer: &targetPlayerIDPtr,
		GuessedRank:  &guessedRank,
	}
	s.lastError = s.engine.ProcessAction(playerIDTyped, finalAction)
	return nil
}

func (s *EngineTestSuite) thenTheGameShouldHavePlayers(expectedCount int) error {
	actualCount := len(s.engine.Game.Players)
	if actualCount != expectedCount {
		return fmt.Errorf("expected %d players, but got %d", expectedCount, actualCount)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldExistInTheGame(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("expected player %s to exist in game", playerID)
	}
	return nil
}

func (s *EngineTestSuite) thenPlayerShouldHaveName(playerID, expectedName string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}
	if player.Name != expectedName {
		return fmt.Errorf("expected player %s to have name '%s', but got '%s'", playerID, expectedName, player.Name)
	}
	return nil
}

func (s *EngineTestSuite) givenPlayerHasCardInTheirHand(playerID, cardName string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return fmt.Errorf("player %s not found", playerID)
	}

	var card model.Card
	switch cardName {
	case "Guard":
		card = model.CardGuard
	case "Priest":
		card = model.CardPriest
	case "Baron":
		card = model.CardBaron
	case "Handmaid":
		card = model.CardHandmaid
	case "Prince":
		card = model.CardPrince
	case "King":
		card = model.CardKing
	case "Countess":
		card = model.CardCountess
	case "Princess":
		card = model.CardPrincess
	default:
		return fmt.Errorf("unknown card type: %s", cardName)
	}

	visibleTo := make(map[model.PlayerID]bool)
	for _, p := range s.engine.Game.Players {
		visibleTo[p.ID] = p.ID == player.ID
	}

	// Replace the hand with just this card
	player.Hand = []model.Redactable[model.Card]{{
		Value:     card,
		VisibleTo: visibleTo,
	}}
	return nil
}

func (s *EngineTestSuite) givenItIsPlayersTurn(playerID string) error {
	for i, player := range s.engine.Game.Players {
		if player.ID == model.PlayerID(playerID) {
			s.engine.Game.CurrentPlayer = i
			return nil
		}
	}
	return fmt.Errorf("player %s not found", playerID)
}

func (s *EngineTestSuite) thenPlayerShouldBeEliminated(playerID string) error {
	player := s.engine.Game.GetPlayer(model.PlayerID(playerID))
	if player == nil {
		return s.wrapError(fmt.Errorf("player %s not found", playerID))
	}
	if !player.IsOut {
		return s.wrapError(fmt.Errorf("expected player %s to be eliminated, but they are not", player.ID))
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
	ctx.Given(`^player ([A-Z]) has card "([^"]*)" in their hand$`, suite.givenPlayerHasCardInTheirHand)
	ctx.Given(`^it is player ([A-Z])'s turn$`, suite.givenItIsPlayersTurn)
	ctx.Given(`^player ([A-Z]) draws a card$`, suite.whenPlayerDrawsACard)

	ctx.When(`^a round starts$`, suite.whenARoundStarts)
	ctx.When(`^player ([A-Z]) draws a card$`, suite.whenPlayerDrawsACard)
	ctx.When(`^player ([A-Z]) performs action "add_player" with new player "([A-Z])" named "([^"]*)"$`, suite.whenPlayerPerformsActionAddPlayer)
	ctx.When(`^player ([A-Z]) performs action "start_game"$`, suite.whenPlayerPerformsActionStartGame)
	ctx.When(`^player ([A-Z]) performs action "play_guard" targeting player ([A-Z]) guessing (\d+)$`, suite.whenPlayerPerformsActionPlayGuardTargetingGuessing)

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
	ctx.Then(`^an error occurs$`, suite.thenAnErrorOccurs)
	ctx.Then(`^no error occurs$`, suite.thenNoErrorOccurs)
	ctx.Then(`^the game should have (\d+) players$`, suite.thenTheGameShouldHavePlayers)
	ctx.Then(`^player ([A-Z]) should exist in the game$`, suite.thenPlayerShouldExistInTheGame)
	ctx.Then(`^player ([A-Z]) should have name "([^"]*)"$`, suite.thenPlayerShouldHaveName)
	ctx.Then(`^player ([A-Z]) should be eliminated$`, suite.thenPlayerShouldBeEliminated)
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
