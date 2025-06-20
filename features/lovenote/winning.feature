Feature: winning conditions in Love Note

  Background:
    Given the game being played was Love note

  Scenario: last player remaining wins the round
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And Alice had the following love note cards in their hand:
      | Guard |
      | King  |
    And Bob had the following love note cards in their hand:
      | Princess |
    And it was Alice's turn
    When Alice plays the love note card Guard against Bob with the guess "Princess"
    Then a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
    And Alice will have 1 point

  Scenario: deck runs out and highest closeness wins
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And it was Alice's turn
    And the love note deck is empty
    And Alice had the following love note cards in their hand:
      | King  |
      | Priest |
    And Bob had the following love note cards in their hand:
      | Baron |
    When Alice plays the love note card King against Bob
    Then a "end-of-round-showdown" event will be raised
    And a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
    And Alice will have 1 point
    And Bob will have 0 points

  Scenario: deck runs out with tied closeness, discard sum breaks tie
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And it was Alice's turn
    And the love note deck is empty
    And Alice had the following love note cards in their hand:
      | Countess |
      | King |
    And Bob had the following love note cards in their hand:
      | King |
    And Alice had the following love note cards in their discard pile:
      | Priest |
      | Priest    |
    And Bob had the following love note cards in their discard pile:
      | Guard  |
      | Guard |
    When Alice plays the love note card Countess
    Then a "end-of-round-showdown" event will be raised
    And a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
    And Alice will have 1 point
    And Bob will have 0 points

  Scenario: deck runs out with complete tie, both players win
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And it was Alice's turn
    And the love note deck is empty
    And Alice had the following love note cards in their hand:
      | Countess |
      | King |
    And Bob had the following love note cards in their hand:
      | King |
    And Alice had the following love note cards in their discard pile:
      | Guard |
      | Guard    |
    And Bob had the following love note cards in their discard pile:
      | Priest |
      | Countess    |
    When Alice plays the love note card Countess
    Then a "end-of-round-showdown" event will be raised
    And a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
      | Bob   |
    And Alice will have 1 point
    And Bob will have 1 point

  Scenario: player wins the overall game by reaching token target
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And Alice had 6 points
    And Bob had 3 points
    And it was Alice's turn
    And the love note deck is empty
    And Alice had the following love note cards in their hand:
      | King |
      | Priest |
    And Bob had the following love note cards in their hand:
      | Baron |
    When Alice plays the love note card King against Bob
    Then a "end-of-round-showdown" event will be raised
    And a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
    And Alice will have 7 points
    And a "game-over" event will be raised
    And the game winners will be:
      | name  |
      | Alice |
    And the game phase will be end

  Scenario: multiple players win the overall game simultaneously
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
    And the game was started by Alice
    And Alice had 4 points
    And Bob had 4 points
    And Charlie had 2 points
    And it was Alice's turn
    And the love note deck is empty
    And Alice had the following love note cards in their discard pile:
      | Guard |
    And Bob had the following love note cards in their discard pile:
      | Countess |
    And Alice had the following love note cards in their hand:
      | King |
      | King |
    And Bob had the following love note cards in their hand:
      | King |
    And Charlie had the following love note cards in their hand:
      | Baron |
    When Alice plays the love note card King against Bob
    Then a "end-of-round-showdown" event will be raised
    And a "round-over" event will be raised
    And the round winners will be:
      | name  |
      | Alice |
      | Bob   |
    And Alice will have 5 points
    And Bob will have 5 points
    And Charlie will have 2 points
    And a "game-over" event will be raised
    And the game winners will be:
      | name  |
      | Alice |
      | Bob   |
    And the game phase will be end