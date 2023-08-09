Feature: playing princess cards

  Background:
    Given the game being played was Love note
    And the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
    And the game was started by Alice
    And the turn order was:
      | Alice   |
      | Bob     |
    And it was Alice's turn
    And Alice had the following love note cards in their hand:
      | Princess |
      | Guard  |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice uses a princess
    When Alice plays the love note card Princess
    Then it will be Bob's turn
    And Bob will have 1 points
