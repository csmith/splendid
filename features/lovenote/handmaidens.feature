Feature: the game is set up properly

  Background:
    Given the game being played was Love note
    And the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
    And the game was started by Alice
    And the turn order was:
      | Alice   |
      | Bob     |
      | Charlie |
    And it was Alice's turn
    And Alice had the following love note cards in their hand:
      | Handmaid |
      | Guard  |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice uses a handmaiden
    When Alice plays the love note card Handmaid
    Then Alice will be protected by a love note handmaid
