Feature: the game is set up properly

  Background:
    Given the game being played was Love note

  Scenario: A game is started with one player
    Given the following players joined the game:
      | name  |
      | Alice |
    And the game was started by Alice
    Then an "Action start not available" error will occur
    And the game phase will be setup

  Scenario: A game is started with two players
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    Then the game phase will be play
    And there will be 3 cards in the love note unused deck
    And there will be 1 card in the love note spare deck
    And the current player will have 2 love note cards in their hand
    And the winner will need 7 love note points to win

  Scenario: A game is started with three players
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
    And the game was started by Alice
    Then the game phase will be play
    And there will be 0 cards in the love note unused deck
    And there will be 1 card in the love note spare deck
    And the current player will have 2 love note cards in their hand
    And the winner will need 5 love note points to win


  Scenario: A game is started with four players
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
      | David   |
    And the game was started by Alice
    Then the game phase will be play
    And there will be 0 cards in the love note unused deck
    And there will be 1 card in the love note spare deck
    And the current player will have 2 love note cards in their hand
    And the winner will need 4 love note points to win

  Scenario: A game is started with five players
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
      | David   |
      | Eve     |
    And the game was started by Alice
    Then an "Action start not available" error will occur
    And the game phase will be setup