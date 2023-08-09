Feature: the game is set up properly

  Background:
    Given the game being played was Splendid

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
    And the following splendid tokens will be available:
      | type     | amount |
      | emerald  | 4      |
      | sapphire | 4      |
      | ruby     | 4      |
      | diamond  | 4      |
      | onyx     | 4      |
      | gold     | 5      |
    And there will be 3 splendid nobles available

  Scenario: A game is started with three players
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
    And the game was started by Alice
    Then the game phase will be play
    And the following splendid tokens will be available:
      | type     | amount |
      | emerald  | 5      |
      | sapphire | 5      |
      | ruby     | 5      |
      | diamond  | 5      |
      | onyx     | 5      |
      | gold     | 5      |
    And there will be 4 splendid nobles available

  Scenario: A game is started with four players
    Given the following players joined the game:
      | name    |
      | Alice   |
      | Bob     |
      | Charlie |
      | David   |
    And the game was started by Alice
    Then the game phase will be play
    And the following splendid tokens will be available:
      | type     | amount |
      | emerald  | 7      |
      | sapphire | 7      |
      | ruby     | 7      |
      | diamond  | 7      |
      | onyx     | 7      |
      | gold     | 5      |
    And there will be 5 splendid nobles available

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