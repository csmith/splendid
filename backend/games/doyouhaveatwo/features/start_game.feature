Feature: Start Game
  Background:
    Given a game with 3 players
    And the game has phase "setup"

  Scenario: Player is able to perform action start_game when game is in setup phase
    When player A performs action "start_game"
    Then no error occurs
    And the game phase should be "draw"
    And the round number should be 1
    And player A should have exactly 1 card in their hand
    And player B should have exactly 1 card in their hand
    And player C should have exactly 1 card in their hand

  Scenario: Player cannot perform action start_game when game is not in setup phase
    Given the game has phase "play"
    When player A performs action "start_game"
    Then an error occurs

  Scenario: Starting game with minimum players works
    Given a game with 2 players
    When player A performs action "start_game"
    Then no error occurs
    And the game phase should be "draw"
    And the round number should be 1
    And player A should have exactly 1 card in their hand
    And player B should have exactly 1 card in their hand

  Scenario: Starting game with maximum players works
    Given a game with 4 players
    When player A performs action "start_game"
    Then no error occurs
    And the game phase should be "draw"
    And the round number should be 1
    And player A should have exactly 1 card in their hand
    And player B should have exactly 1 card in their hand
    And player C should have exactly 1 card in their hand
    And player D should have exactly 1 card in their hand

  Scenario: Cannot start game with insufficient players
    Given a game with 1 players
    When player A performs action "start_game"
    Then an error occurs

  Scenario: Start game initializes deck correctly
    When player A performs action "start_game"
    Then no error occurs
    And there are 5 Guard cards in the game
    And there are 2 Priest cards in the game
    And there are 2 Baron cards in the game
    And there are 2 Handmaid cards in the game
    And there are 2 Prince cards in the game
    And there are 1 King cards in the game
    And there are 1 Countess cards in the game
    And there are 1 Princess cards in the game
    And a card should be removed from the game