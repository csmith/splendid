Feature: playing prince cards

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
      | Prince |
      | Guard  |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice tries to use a prince against a protected player
    Given Bob is protected by a love note handmaid
    When Alice plays the love note card Prince against Bob
    Then a "Player Bob cannot be targeted because they have handmaiden cover" error will occur
    And Bob will not be eliminated
    And Bob will have the following love note cards in their hand:
      | Priest |
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | Prince |
      | Guard  |

  Scenario: Alice uses a prince against Bob
    Given the next love note card in the deck is Princess
    And the next love note card in the deck is Guard
    When Alice plays the love note card Prince against Bob
    Then Bob will have the following love note cards in their hand:
      | Guard    |
      | Princess |
    And Bob will have the following love note cards in their discards:
      | Priest |
    And Alice will have the following love note cards in their hand:
      | Guard |
    And Alice will have the following love note cards in their discards:
      | Prince |

  Scenario: Alice uses a prince to make Bob discard the princess
    Given Bob had the following love note cards in their hand:
      | Princess |
    When Alice plays the love note card Prince against Bob
    Then Bob will be eliminated
    And Bob will have the following love note cards in their discards:
      | Princess |
    And Alice will have the following love note cards in their hand:
      | Guard |
    And Alice will have the following love note cards in their discards:
      | Prince |
