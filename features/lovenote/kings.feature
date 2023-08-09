Feature: playing king cards

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
      | King  |
      | Guard |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice tries to use a king against a protected player
    Given Bob is protected by a love note handmaid
    When Alice plays the love note card King against Bob
    Then a "Player Bob cannot be targeted because they have handmaiden cover" error will occur
    And Bob will not be eliminated
    And Bob will have the following love note cards in their hand:
      | Priest |
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | King  |
      | Guard |

  Scenario: Alice uses a king against Bob
    Given the next love note card in the deck is Prince
    When Alice plays the love note card King against Bob
    Then Bob will have the following love note cards in their hand:
      | Guard  |
      | Prince |
    And Alice will have the following love note cards in their hand:
      | Priest |
    And Alice will have the following love note cards in their discards:
      | King |
