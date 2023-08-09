Feature: playing barons cards

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
      | Baron |
      | Baron |
    And Bob had the following love note cards in their hand:
      | Priest |
    And Charlie had the following love note cards in their hand:
      | Princess |

  Scenario: Alice tries to use a baron against a protected player
    Given Bob is protected by a love note handmaid
    When Alice plays the love note card Baron against Bob
    Then a "Player Bob cannot be targeted because they have handmaiden cover" error will occur
    And Bob will not be eliminated
    And a "hand-revealed" event will be not raised
    And it will be Alice's turn still
    And the game phase will be play still
    And Alice will have the following love note cards in their hand:
      | Baron |
      | Baron |

  Scenario: Alice uses a baron when all other players are protected
    Given Bob is protected by a love note handmaid
    And Charlie is protected by a love note handmaid
    When Alice plays the love note card Baron
    Then Alice will not be eliminated
    And Bob will not be eliminated
    And Charlie will not be eliminated
    And a "card-no-op" event will be raised

  Scenario: Alice uses a Baron against Bob and he is knocked out
    When Alice plays the love note card Baron against Bob
    Then a "hand-revealed" event will be raised
    And Bob will be eliminated
    And Alice will not be eliminated

  Scenario: Alice uses a Baron against Charlie and she is knocked out
    When Alice plays the love note card Baron against Charlie
    Then a "hand-revealed" event will be raised
    And Alice will be eliminated
    And Charlie will not be eliminated