Feature: players can take tokens

  Background:
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice

  Scenario: Alice takes two tokens of the same type from a stack of four
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type    | amount |
      | emerald | 4      |
    When Alice draws the following splendid tokens:
      | type    | amount |
      | emerald | 2      |
    Then Alice will have the following splendid tokens:
      | type    | amount |
      | emerald | 2      |
    And the following splendid tokens will be available:
      | type    | amount |
      | emerald | 2      |
    And it will be Bob's turn

  Scenario: Alice takes two tokens of the same type from a stack of three
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type    | amount |
      | emerald | 3      |
    When Alice draws the following splendid tokens:
      | type    | amount |
      | emerald | 2      |
    Then a "Can't draw double emerald as only 3 are available" error will occur
    Then Alice will have the following splendid tokens:
      | type    | amount |
      | emerald | 0      |
    And the following splendid tokens will be available:
      | type    | amount |
      | emerald | 3      |
    And it will be Alice's turn still

  Scenario: Alice takes three tokens of different types
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type     | amount |
      | emerald  | 3      |
      | ruby     | 3      |
      | sapphire | 3      |
    When Alice draws the following splendid tokens:
      | type     | amount |
      | emerald  | 1      |
      | ruby     | 1      |
      | sapphire | 1      |
    Then Alice will have the following splendid tokens:
      | type     | amount |
      | emerald  | 1      |
      | ruby     | 1      |
      | sapphire | 1      |
    And the following splendid tokens will be available:
      | type     | amount |
      | emerald  | 2      |
      | ruby     | 2      |
      | sapphire | 2      |
    And it will be Bob's turn

  Scenario: Alice takes more than three tokens of different types
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type     | amount |
      | emerald  | 3      |
      | ruby     | 3      |
      | sapphire | 3      |
      | onyx     | 3      |
    When Alice draws the following splendid tokens:
      | type     | amount |
      | emerald  | 1      |
      | ruby     | 1      |
      | sapphire | 1      |
      | onyx     | 1      |
    Then a "You must draw 3 tokens of different types, or 2 of the same type" error will occur
    Then Alice will have the following splendid tokens:
      | type     | amount |
      | emerald  | 0      |
      | ruby     | 0      |
      | sapphire | 0      |
      | onyx     | 0      |
    And the following splendid tokens will be available:
      | type     | amount |
      | emerald  | 3      |
      | ruby     | 3      |
      | sapphire | 3      |
      | onyx     | 3      |
    And it will be Alice's turn still

  Scenario: Alice takes more than two tokens of the same type
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type    | amount |
      | emerald | 5      |
    When Alice draws the following splendid tokens:
      | type    | amount |
      | emerald | 3      |
    Then a "Can't draw more than 1 emerald if drawing 3 total" error will occur
    Then Alice will have the following splendid tokens:
      | type    | amount |
      | emerald | 0      |
    And the following splendid tokens will be available:
      | type    | amount |
      | emerald | 5      |
    And it will be Alice's turn still

  Scenario: Alice takes two tokens of one type, and one of another
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type    | amount |
      | emerald | 5      |
      | onyx    | 1      |
    When Alice draws the following splendid tokens:
      | type    | amount |
      | emerald | 2      |
      | onyx    | 1      |
    Then a "You must draw 3 tokens of different types, or 2 of the same type" error will occur
    Then Alice will have the following splendid tokens:
      | type    | amount |
      | emerald | 0      |
      | onyx    | 0      |
    And the following splendid tokens will be available:
      | type    | amount |
      | emerald | 5      |
      | onyx    | 1      |
    And it will be Alice's turn still

  Scenario: Alice takes gold tokens
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type    | amount |
      | emerald | 5      |
      | gold    | 5      |
    When Alice draws the following splendid tokens:
      | type | amount |
      | gold | 2      |
    Then a "You must draw 3 tokens of different types, or 2 of the same type" error will occur
    Then Alice will have the following splendid tokens:
      | type | amount |
      | gold | 0      |
    And the following splendid tokens will be available:
      | type    | amount |
      | emerald | 5      |
      | gold    | 5      |
    And it will be Alice's turn still

  Scenario: Alice takes tokens that aren't available
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type     | amount |
      | onyx     | 0      |
      | sapphire | 1      |
      | ruby     | 1      |
      | diamond  | 1      |
    When Alice draws the following splendid tokens:
      | type     | amount |
      | onyx     | 1      |
      | sapphire | 1      |
      | ruby     | 1      |
    Then a "Can't draw onyx as none are available" error will occur
    Then Alice will have the following splendid tokens:
      | type     | amount |
      | onyx     | 0      |
      | sapphire | 0      |
      | ruby     | 0      |
    And the following splendid tokens will be available:
      | type     | amount |
      | onyx     | 0      |
      | sapphire | 1      |
      | ruby     | 1      |
      | diamond  | 1      |
    And it will be Alice's turn still

  Scenario: Alice takes two different tokens because no more are available
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type     | amount |
      | sapphire | 1      |
      | ruby     | 1      |
    When Alice draws the following splendid tokens:
      | type     | amount |
      | sapphire | 1      |
      | ruby     | 1      |
    Then Alice will have the following splendid tokens:
      | type     | amount |
      | sapphire | 1      |
      | ruby     | 1      |
    And the following splendid tokens will be available:
      | type     | amount |
      | sapphire | 0      |
      | ruby     | 0      |
    And it will be Bob's turn

  Scenario: Alice takes tokens and ends up with more than 10
    Given it was Alice's turn
    And the following splendid tokens were available:
      | type | amount |
      | onyx | 4      |
    And Alice had the following splendid tokens:
      | type | amount |
      | ruby | 9      |
    When Alice draws the following splendid tokens:
      | type | amount |
      | onyx | 2      |
    Then Alice will have the following splendid tokens:
      | type | amount |
      | onyx | 2      |
      | ruby | 9      |
    And the game phase will be discard
    And it will be Alice's turn still

  Scenario: Bob tries to take tokens on Alice's turn
    Given it was Alice's turn
    When Bob draws the following splendid tokens:
      | type | amount |
      | onyx | 2      |
    Then an "Action take-tokens not available" error will occur
    And it will be Alice's turn still