Feature: Round Win and New Round Start
  Background:
    Given a game with 3 players
    And player A has the following cards in their hand:
      | Princess |
    And player B has the following cards in their hand:
      | King |
    And player C has the following cards in their hand:
      | Guard |
    And the game has phase "play"
    And it is player A's turn

  Scenario: Single player wins round and gets a token
    Given player B is eliminated
    And player C is eliminated
    When the round ends with winners A
    Then player A should have 1 tokens
    And the following event occurred: "player_token_awarded"
    And the following event occurred: "last_round_winners_updated"
    And the round number should be 1
    And the game phase should be "draw"
    And player A should not be eliminated
    And player B should not be eliminated
    And player C should not be eliminated
    And player A should have exactly 1 card in their hand
    And player B should have exactly 1 card in their hand
    And player C should have exactly 1 card in their hand

  Scenario: Multiple players tie and both get tokens
    Given player C is eliminated
    And player A has the following cards in their hand:
      | King |
    And player B has the following cards in their hand:
      | King |
    When a showdown occurs
    Then players A,B should win the round
    And player A should have 1 tokens
    And player B should have 1 tokens
    And player C should have 0 tokens
    And the round number should be 1

  Scenario: Showdown determines winner based on card values
    Given player A has the following cards in their hand:
      | Princess |
    And player B has the following cards in their hand:
      | King |
    And player C has the following cards in their hand:
      | Guard |
    When a showdown occurs
    Then players A should win the round
    And player A should have 1 tokens
    And player B should have 0 tokens
    And player C should have 0 tokens

  Scenario: Player wins game when reaching token threshold
    Given player A has 4 tokens
    And player B is eliminated
    And player C is eliminated
    When the round ends with winners A
    Then player A should have 5 tokens
    And the game should end
    And player A should win the game
    And the game phase should be "game_end"

  Scenario: Round winners start next round first
    Given player B is eliminated
    And player C is eliminated
    And player A won the last round
    When a round starts
    Then it should be player A's turn
    And the round number should be incremented

  Scenario: New round resets player states correctly
    Given player B is eliminated
    And player A is protected
    And player A has 2 cards in discard pile
    And player B has 1 cards in discard pile
    When a round starts
    Then player A should not be eliminated
    And player B should not be eliminated
    And player A should not be protected
    And player B should not be protected
    And player A should have 0 cards in discard pile
    And player B should have 0 cards in discard pile
    And the deck should have 12 cards remaining
    And a card should be removed from the game