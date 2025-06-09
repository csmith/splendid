Feature: Play Countess

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Countess |
      | Guard    |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Player is able to perform action play_countess when it's their turn and they have a Countess
    Given it is player A's turn
    When player A sends action {"type": "play_countess", "player": "A", "card_name": "Countess"}
    Then no error occurs
    And player A should have 1 cards in discard pile

  Scenario: Countess must be played when holding King
    Given it is player A's turn
    And player A has the following cards in their hand:
      | King     |
      | Countess |
    Then the available actions should be:
      | player | action                  |
      | A      | play_countess(player=A) |

  Scenario: Countess must be played when holding Prince
    Given it is player A's turn
    And player A has the following cards in their hand:
      | Prince   |
      | Countess |
    Then the available actions should be:
      | player | action                  |
      | A      | play_countess(player=A) |

  Scenario: Countess can be played voluntarily when not holding King or Prince
    Given it is player A's turn
    And player A has the following cards in their hand:
      | Baron    |
      | Countess |
    Then the available actions should be:
      | player | action                  |
      | A      | play_countess(player=A) |
      | A      | play_baron(player=A)    |

  Scenario: Player cannot play King when holding Countess
    Given it is player A's turn
    And player A has the following cards in their hand:
      | King     |
      | Countess |
    When player A sends action {"type": "play_king", "player": "A", "card_name": "King"}
    Then an error occurs

  Scenario: Player cannot play Prince when holding Countess
    Given it is player A's turn
    And player A has the following cards in their hand:
      | Prince   |
      | Countess |
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then an error occurs

  Scenario: Player cannot perform action play_countess when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_countess", "player": "A", "card_name": "Countess"}
    Then an error occurs

  Scenario: Player cannot perform action play_countess when they don't have a Countess
    Given it is player B's turn
    When player B sends action {"type": "play_countess", "player": "B", "card_name": "Countess"}
    Then an error occurs

  Scenario: Available actions when player holds Countess with a lesser card
    Given it is player A's turn
    Then the available actions should be:
      | player | action                  |
      | A      | play_countess(player=A) |
      | A      | play_guard(player=A)    |

  Scenario: Countess action is immediately complete
    Given it is player A's turn
    When player A sends action {"type": "play_countess", "player": "A", "card_name": "Countess"}
    Then no error occurs
    And it should be player B's turn