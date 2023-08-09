Feature: playing guard cards

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
      | Guard |
      | King |
    And Bob had the following love note cards in their hand:
      | Priest |

  Scenario: Alice tries to use a guard to guess an invalid card
    When Alice plays the love note card Guard against Bob with the guess "bloop"
    Then an "Invalid guessed type bloop" error will occur
    And Bob will not be eliminated
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | Guard |
      | King  |

  Scenario: Alice tries to use a guard against a protected player
    Given Bob is protected by a love note handmaid
    When Alice plays the love note card Guard against Bob with the guess "Priest"
    Then a "Player Bob cannot be targeted because they have handmaiden cover" error will occur
    And Bob will not be eliminated
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | Guard |
      | King  |

  Scenario: Alice uses a guard when all other players are protected
    Given Bob is protected by a love note handmaid
    And Charlie is protected by a love note handmaid
    When Alice plays the love note card Guard
    Then Alice will not be eliminated
    And Bob will not be eliminated
    And Charlie will not be eliminated
    And a "card-no-op" event will be raised
    And Alice will have the following love note cards in their hand:
      | King  |
    And Alice will have the following love note cards in their discards:
      | Guard |

  Scenario: Alice uses a guard to guess an incorrect card
    When Alice plays the love note card Guard against Bob with the guess "Princess"
    Then Bob will not be eliminated
    And a "guard-failed" event will be raised
    And Alice will have the following love note cards in their hand:
      | King  |
    And Alice will have the following love note cards in their discards:
      | Guard |

  Scenario: Alice uses a guard to guess a correct card
    When Alice plays the love note card Guard against Bob with the guess "Priest"
    Then Bob will be eliminated
    And Alice will have the following love note cards in their hand:
      | King  |
    And Alice will have the following love note cards in their discards:
      | Guard |
    And Bob will have the following love note cards in their discards:
      | Priest |