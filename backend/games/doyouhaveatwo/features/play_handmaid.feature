Feature: Play Handmaid

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Handmaid |
      | Guard    |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Player is able to perform action play_handmaid when it's their turn and they have a Handmaid
    Given it is player A's turn
    When player A sends action {"type": "play_handmaid", "player": "A", "card_name": "Handmaid"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player A should be protected

  Scenario: Player cannot perform action play_handmaid when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_handmaid", "player": "A", "card_name": "Handmaid"}
    Then an error occurs

  Scenario: Player cannot perform action play_handmaid when they don't have a Handmaid
    Given it is player B's turn
    When player B sends action {"type": "play_handmaid", "player": "B", "card_name": "Handmaid"}
    Then an error occurs

  Scenario: Available actions when player holds Handmaid
    Given it is player A's turn
    Then the available actions should be:
      | player | action                  |
      | A      | play_handmaid(player=A) |
      | A      | play_guard(player=A)    |

  Scenario: Handmaid action is immediately complete
    Given it is player A's turn
    When player A sends action {"type": "play_handmaid", "player": "A", "card_name": "Handmaid"}
    Then no error occurs
    And it should be player B's turn