Feature: playing priest cards

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
      | Priest |
      | Guard  |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice tries to use a priest against a protected player
    Given Bob is protected by a love note handmaid
    When Alice plays the love note card Priest against Bob
    Then a "Player Bob cannot be targeted because they have handmaiden cover" error will occur
    And Bob will not be eliminated
    And a "hand-revealed" event will be not raised
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | Priest |
      | Guard  |

  Scenario: Alice uses a priest
    When Alice plays the love note card Priest against Bob
    Then a "hand-revealed" event will be raised
