Feature: playing countess cards

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
      | Countess |
      | Guard    |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice uses a countess
    When Alice plays the love note card Countess
    Then Alice will have the following love note cards in their hand:
      | Guard |
    And Alice will have the following love note cards in their discards:
      | Countess |

  Scenario: Alice uses a countess when she has a king
    Given Alice had the following love note cards in their hand:
      | Countess |
      | King     |
    When Alice plays the love note card Countess
    Then Alice will have the following love note cards in their hand:
      | King |
    And Alice will have the following love note cards in their discards:
      | Countess |

  Scenario: Alice tries to not use a countess when she has a king
    Given Alice had the following love note cards in their hand:
      | Countess |
      | King     |
    When Alice plays the love note card King
    Then a "You must discard the Countess" error will occur
    Then Alice will have the following love note cards in their hand:
      | Countess |
      | King |
    And it will be Alice's turn still

  Scenario: Alice uses a countess when she has a prince
    Given Alice had the following love note cards in their hand:
      | Countess |
      | Prince   |
    When Alice plays the love note card Countess
    Then Alice will have the following love note cards in their hand:
      | Prince |
    And Alice will have the following love note cards in their discards:
      | Countess |

  Scenario: Alice tries to not use a countess when she has a prince
    Given Alice had the following love note cards in their hand:
      | Countess |
      | Prince     |
    When Alice plays the love note card Prince
    Then a "You must discard the Countess" error will occur
    Then Alice will have the following love note cards in their hand:
      | Countess |
      | Prince |
    And it will be Alice's turn still