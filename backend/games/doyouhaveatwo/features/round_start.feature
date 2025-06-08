Feature: Round Start
  Background:
    Given a game with 3 players
    And the game has phase "setup"
    And the current round is 2

  Scenario: Round number increments on start
    When a round starts
    Then the round number should be 3

  Scenario: Deck is created with correct cards
    When a round starts
    Then there are 5 Guard cards in the game
    And there are 2 Priest cards in the game
    And there are 2 Baron cards in the game
    And there are 2 Handmaid cards in the game
    And there are 2 Prince cards in the game
    And there are 1 King cards in the game
    And there are 1 Countess cards in the game
    And there are 1 Princess cards in the game

  Scenario: One card is removed from deck
    When a round starts
    Then a card should be removed from the game
    And the removed card should not be visible to any player

  Scenario: Each player receives one card
    When a round starts
    Then player A should have exactly 1 card in their hand
    And player B should have exactly 1 card in their hand
    And player C should have exactly 1 card in their hand
    And player A's cards should only be visible to themselves
    And player B's cards should only be visible to themselves
    And player C's cards should only be visible to themselves
    And the deck should have 12 cards remaining

  Scenario: Player states are reset
    Given player A has 2 cards in discard pile
    And player B has 2 cards in discard pile
    And player C has 2 cards in discard pile
    And player A is eliminated
    And player B is protected
    When a round starts
    Then player A should not be eliminated
    And player B should not be eliminated
    And player C should not be eliminated
    And player A should not be protected
    And player B should not be protected
    And player C should not be protected
    And player A should have 0 cards in discard pile
    And player B should have 0 cards in discard pile
    And player C should have 0 cards in discard pile

  Scenario: Game phase is set to draw
    When a round starts
    Then the game phase should be "draw"
