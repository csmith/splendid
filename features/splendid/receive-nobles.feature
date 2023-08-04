Feature: players can receive nobles

  Background:
    Given the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And it was Alice's turn
    And the following cards were visible:
      # Level / Points / Bonus / Cost: E-S-R-D-O
      | 3/3/emerald/07200 | 3/4/sapphire/50500 | 3/4/emerald/07300 | 3/4/onyx/44400    |
      | 2/1/onyx/01200    | 2/1/sapphire/00022 | 2/2/ruby/20202    | 2/1/emerald/00220 |
      | 1/0/onyx/10100    | 1/0/sapphire/00011 | 1/1/ruby/10101    | 1/0/emerald/00110 |
    And the top card of deck 3 was 3/3/diamond/07000
    And the top card of deck 2 was 2/2/diamond/20200
    And the top card of deck 1 was 1/1/diamond/10100
    And the following nobles were available:
      | 03300 |
      | 00033 |
      | 22222 |

  Scenario: Alice buys a card which enables a noble visit
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 2      |
    And Alice had the following tokens:
      | type    | amount |
      | emerald | 1      |
      | ruby    | 1      |
      | onyx    | 1      |
    When Alice buys the card 1/1/ruby/10101
    Then the game phase will be noble
    And it will be Alice's turn still

  Scenario: Alice takes some tokens while she has an eligible noble
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 3      |
    And the following tokens were available:
      | type    | amount |
      | emerald | 1      |
      | ruby    | 1      |
      | onyx    | 1      |
    When Alice draws the following tokens:
      | type    | amount |
      | emerald | 1      |
      | ruby    | 1      |
      | onyx    | 1      |
    Then the game phase will be noble
    And it will be Alice's turn still

  Scenario: Alice tries to receive a noble at the start of a turn
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 3      |
    When Alice receives the noble 03300
    Then an "Action receive-noble not found" error will occur
    And the game phase will be play
    And it will be Alice's turn still

  Scenario: Alice tries to receive an unknown noble
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 3      |
    And the game phase was noble
    When Alice receives the noble 99999
    Then a "Noble not found" error will occur
    And the game phase will be noble
    And it will be Alice's turn still

  Scenario: Alice tries to receive an ineligible noble
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 3      |
    And the game phase was noble
    When Alice receives the noble 22222
    Then a "Not eligible for noble" error will occur
    And the game phase will be noble
    And it will be Alice's turn still

  Scenario: Alice receives a noble
    Given Alice had the following bonuses:
      | type     | amount |
      | sapphire | 3      |
      | ruby     | 3      |
    And the game phase was noble
    When Alice receives the noble 03300
    Then Alice will have 3 points
    And the game phase will be play
    And it will be Bob's turn