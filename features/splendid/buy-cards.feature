Feature: players can buy cards

  Background:
    Given the game being played was Splendid
    And the following players joined the game:
      | name  |
      | Alice |
      | Bob   |
    And the game was started by Alice
    And it was Alice's turn
    And the following splendid cards were visible:
      # Level / Points / Bonus / Cost: E-S-R-D-O
      | 3/3/emerald/07200 | 3/4/sapphire/50500 | 3/4/emerald/07300 | 3/4/onyx/44400    |
      | 2/1/onyx/01200    | 2/1/sapphire/00022 | 2/2/ruby/20202    | 2/1/emerald/00220 |
      | 1/0/onyx/10100    | 1/0/sapphire/00011 | 1/1/ruby/10101    | 1/0/emerald/00110 |
    And the top splendid card of deck 3 was 3/3/diamond/07000
    And the top splendid card of deck 2 was 2/2/diamond/20200
    And the top splendid card of deck 1 was 1/1/diamond/10100

  Scenario: Alice tries to buy a card she can't afford
    When Alice buys the splendid card 3/3/emerald/07200
    Then a "Cannot afford card" error will occur
    And it will be Alice's turn still
    And the splendid card in row 1 column 1 will be 3/3/emerald/07200
    And Alice will have 0 points

  Scenario: Alice tries to buy a card that doesn't exist
    When Alice buys the splendid card 7/7/nonsense/99999
    Then a "Card not found" error will occur
    And it will be Alice's turn still
    And Alice will have 0 points

  Scenario: Alice buys card she can afford with normal tokens
    Given Alice had the following splendid tokens:
      | type    | amount |
      | diamond | 1      |
      | onyx    | 1      |
    When Alice buys the splendid card 1/0/sapphire/00011
    Then Alice will have the following splendid bonuses:
      | type     | amount |
      | sapphire | 1      |
    And Alice will have the following splendid tokens:
      | type    | amount |
      | diamond | 0      |
      | onyx    | 0      |
    And Alice will have 0 points
    And the splendid card in row 3 column 2 will be 1/1/diamond/10100
    And it will be Bob's turn
    And the following splendid tokens will be available:
      | type     | amount |
      | diamond  | 5      |
      | sapphire | 4      |
      | emerald  | 4      |
      | ruby     | 4      |
      | onyx     | 5      |
      | gold     | 5      |

  Scenario: Alice buys card she can afford with gold
    Given Alice had the following splendid tokens:
      | type    | amount |
      | diamond | 1      |
      | gold    | 1      |
    When Alice buys the splendid card 1/0/sapphire/00011
    Then Alice will have the following splendid bonuses:
      | type     | amount |
      | sapphire | 1      |
    And Alice will have the following splendid tokens:
      | type    | amount |
      | diamond | 0      |
      | gold    | 0      |
    And Alice will have 0 points
    And the splendid card in row 3 column 2 will be 1/1/diamond/10100
    And it will be Bob's turn
    And the following splendid tokens will be available:
      | type     | amount |
      | diamond  | 5      |
      | sapphire | 4      |
      | emerald  | 4      |
      | ruby     | 4      |
      | onyx     | 4      |
      | gold     | 6      |

  Scenario: Alice buys card she can afford with bonuses
    Given Alice had the following splendid tokens:
      | type | amount |
      | onyx | 1      |
    And Alice had the following splendid bonuses:
      | type    | amount |
      | diamond | 1      |
    When Alice buys the splendid card 1/0/sapphire/00011
    Then Alice will have the following splendid bonuses:
      | type     | amount |
      | sapphire | 1      |
      | diamond  | 1      |
    And Alice will have the following splendid tokens:
      | type    | amount |
      | diamond | 0      |
      | onyx    | 0      |
    And Alice will have 0 points
    And the splendid card in row 3 column 2 will be 1/1/diamond/10100
    And it will be Bob's turn
    And the following splendid tokens will be available:
      | type     | amount |
      | diamond  | 4      |
      | sapphire | 4      |
      | emerald  | 4      |
      | ruby     | 4      |
      | onyx     | 5      |
      | gold     | 5      |

  Scenario: Alice buys a card that gives points
    Given Alice had the following splendid tokens:
      | type | amount |
      | ruby | 2      |
    And Alice had the following splendid bonuses:
      | type     | amount |
      | sapphire | 7      |
    When Alice buys the splendid card 3/3/emerald/07200
    Then Alice will have the following splendid bonuses:
      | type     | amount |
      | sapphire | 7      |
      | emerald  | 1      |
    And Alice will have the following splendid tokens:
      | type     | amount |
      | ruby     | 0      |
      | sapphire | 0      |
    And Alice will have 3 points
    And the splendid card in row 1 column 1 will be 3/3/diamond/07000
    And it will be Bob's turn
    And the following splendid tokens will be available:
      | type     | amount |
      | diamond  | 4      |
      | sapphire | 4      |
      | emerald  | 4      |
      | ruby     | 6      |
      | onyx     | 4      |
      | gold     | 5      |

  Scenario: Alice buys a card that takes her to 15 points, but she's not the last player
    Given Alice had the following splendid tokens:
      | type | amount |
      | ruby | 2      |
    And Alice had the following splendid bonuses:
      | type     | amount |
      | sapphire | 7      |
    And Alice had 12 points
    And the turn order was:
      | Alice |
      | Bob   |
    When Alice buys the splendid card 3/3/emerald/07200
    Then Alice will have 15 points
    And it will be Bob's turn
    And the game phase will be play
    And this will be the final round

  Scenario: Alice buys a card that takes her to 15 points, and she is the last player
    Given Alice had the following splendid tokens:
      | type | amount |
      | ruby | 2      |
    And Alice had the following splendid bonuses:
      | type     | amount |
      | sapphire | 7      |
    And Alice had 12 points
    And the turn order was:
      | Bob   |
      | Alice |
    When Alice buys the splendid card 3/3/emerald/07200
    Then Alice will have 15 points
    And the game phase will be end

  Scenario: The last player to go in the final round buys a card
    Given Alice had the following splendid tokens:
      | type | amount |
      | ruby | 2      |
    And Alice had the following splendid bonuses:
      | type     | amount |
      | sapphire | 7      |
    And the turn order was:
      | Bob   |
      | Alice |
    And this was the final round
    When Alice buys the splendid card 3/3/emerald/07200
    Then the game phase will be end