Feature: Play Princess

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Princess |
      | Guard    |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Playing Princess eliminates the player immediately
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_no_target", "player": "A", "card_name": "Princess"}
    Then no error occurs
    And player A should have 2 cards in discard pile
    And player A should be eliminated

  Scenario: Player cannot perform action play_princess when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "dyhat:a:play_card_no_target", "player": "A", "card_name": "Princess"}
    Then an error occurs

  Scenario: Player cannot perform action play_princess when they don't have a Princess
    Given it is player B's turn
    When player B sends action {"type": "dyhat:a:play_card_no_target", "player": "B", "card_name": "Princess"}
    Then an error occurs

  Scenario: Princess is discarded when forced by Prince
    Given it is player B's turn
    And player B has the following cards in their hand:
      | Prince |
      | King   |
    And player A has the following cards in their hand:
      | Princess |
    When player B sends action {"type": "dyhat:a:play_card_target_any", "player": "B", "card_name": "Prince"}
    And player B sends action {"type": "dyhat:a:play_card_target_any", "player": "B", "card_name": "Prince", "target_player": "A"}
    Then no error occurs
    And player A should be eliminated
    And player A should have 1 cards in discard pile

  Scenario: Princess must be played when it's the only card
    Given it is player A's turn
    And player A has the following cards in their hand:
      | Princess |
    Then the available actions should be:
      | player | action                  |
      | A      | play_princess(player=A) |

  Scenario: Available actions when player holds Princess with other cards
    Given it is player A's turn
    And player A has the following cards in their hand:
      | Baron    |
      | Princess |
    Then the available actions should be:
      | player | action                  |
      | A      | play_baron(player=A)    |
      | A      | play_princess(player=A) |

  Scenario: Princess action is immediately complete and eliminates player
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_no_target", "player": "A", "card_name": "Princess"}
    Then no error occurs
    And player A should be eliminated
    And it should be player B's turn