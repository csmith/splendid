Feature: players can discard tokens

  Background:
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And the following tokens were available:
      | type     | amount |
      | emerald  | 1      |
      | sapphire | 2      |
      | ruby     | 3      |
      | diamond  | 4      |
      | onyx     | 5      |
      | gold     | 6      |
    And it was Alice's turn
    And Alice had the following tokens:
      | type     | amount |
      | emerald  | 3      |
      | sapphire | 2      |
      | ruby     | 1      |
      | diamond  | 4      |
      | onyx     | 2      |
      | gold     | 1      |
    And the game phase was discard

  Scenario: Alice discards the right number of tokens
    When Alice discards the following tokens:
      | type     | amount |
      | emerald  | 2      |
      | sapphire | 1      |
    Then Alice will have the following tokens:
      | type     | amount |
      | emerald  | 1      |
      | sapphire | 1      |
      | ruby     | 1      |
      | diamond  | 4      |
      | onyx     | 2      |
      | gold     | 1      |
    And the following tokens will be available:
      | type     | amount |
      | emerald  | 3      |
      | sapphire | 3      |
      | ruby     | 3      |
      | diamond  | 4      |
      | onyx     | 5      |
      | gold     | 6      |
    And the game phase will be play
    And it will be Bob's turn

  Scenario: Alice tries to discard too many tokens
    When Alice discards the following tokens:
      | type     | amount |
      | emerald  | 2      |
      | sapphire | 2      |
    Then a "You must discard 3 tokens" error will occur
    And the game phase will be discard still
    And it will be Alice's turn still

  Scenario: Alice tries to discard too few tokens
    When Alice discards the following tokens:
      | type     | amount |
      | emerald  | 1      |
      | sapphire | 1      |
    Then a "You must discard 3 tokens" error will occur
    And the game phase will be discard still
    And it will be Alice's turn still

  Scenario: Alice tries to discard tokens she doesn't have
    When Alice discards the following tokens:
      | type | amount |
      | ruby | 3      |
    Then a "You only have 1 ruby tokens" error will occur
    And the game phase will be discard still
    And it will be Alice's turn still