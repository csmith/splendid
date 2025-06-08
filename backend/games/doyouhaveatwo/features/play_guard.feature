Feature: Play Guard
  Background:
    Given a game with 3 players
    And the game has phase "play"
    And a round starts
    And player A has card "Guard" in their hand
    And player B has card "Priest" in their hand
    And player C has card "Baron" in their hand
    And player A draws a card

  Scenario: Player is able to perform action play_guard when it's their turn and they have a Guard
    Given it is player A's turn
    When player A performs action "play_guard" targeting player B guessing 2
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should be eliminated
    And player B should have 1 cards in discard pile

  Scenario: Player cannot perform action play_guard when it's not their turn
    Given it is player B's turn
    When player A performs action "play_guard" targeting player B guessing 2
    Then an error occurs

  Scenario: Player cannot perform action play_guard when they don't have a Guard
    Given it is player B's turn
    When player B performs action "play_guard" targeting player A guessing 2
    Then an error occurs

  Scenario: Play Guard with incorrect guess does not eliminate target
    Given it is player A's turn
    When player A performs action "play_guard" targeting player B guessing 3
    Then no error occurs
    And player B should not be eliminated
    And player A should have 1 cards in discard pile

  Scenario: Play Guard with correct guess eliminates target
    Given it is player A's turn
    When player A performs action "play_guard" targeting player B guessing 2
    Then no error occurs
    And player B should be eliminated
    And player A should have 1 cards in discard pile

  Scenario: Cannot target protected player with Guard
    Given it is player A's turn
    And player B is protected
    When player A performs action "play_guard" targeting player B guessing 2
    Then an error occurs

  Scenario: Cannot target eliminated player with Guard
    Given it is player A's turn
    And player B is eliminated
    When player A performs action "play_guard" targeting player B guessing 2
    Then an error occurs

  Scenario: Cannot target yourself with Guard
    Given it is player A's turn
    When player A performs action "play_guard" targeting player A guessing 2
    Then an error occurs

  Scenario: Cannot guess Guard with Guard
    Given it is player A's turn
    When player A performs action "play_guard" targeting player B guessing 1
    Then an error occurs