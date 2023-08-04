Feature: players can reserve cards

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

  Scenario: Alice tries to reserve a card but she's got three already
    Given Alice had the following reserved cards:
      | 3/3/emerald/07200 |
      | 3/2/emerald/07200 |
      | 3/1/emerald/07200 |
    When Alice reserves the card 3/3/emerald/07200
    Then an "Action reserve-card not available" error will occur
    And it will be Alice's turn still
    And the card in row 1 column 1 will be 3/3/emerald/07200
    And Alice will have 0 points
    And Alice will have the following tokens:
      | type | amount |
      | gold | 0      |

  Scenario: Alice tries to reserve a card that doesn't exist
    When Alice reserves the card 7/7/nonsense/99999
    Then a "Card not found" error will occur
    And it will be Alice's turn still
    And Alice will have 0 points
    And Alice will have the following tokens:
      | type | amount |
      | gold | 0      |

  Scenario: Alice reserves a card while there are gold tokens available
    Given the following tokens were available:
      | type | amount |
      | gold | 3      |
    When Alice reserves the card 1/0/sapphire/00011
    Then Alice will have the following tokens:
      | type | amount |
      | gold | 1      |
    And Alice will have the following reserved cards:
      | 1/0/sapphire/00011 |
    And Alice will have 0 points
    And the card in row 3 column 2 will be 1/1/diamond/10100
    And it will be Bob's turn
    And the following tokens will be available:
      | type | amount |
      | gold | 2      |

  Scenario: Alice reserves a card while there are no gold tokens available
    Given the following tokens were available:
      | type | amount |
      | gold | 0      |
    When Alice reserves the card 1/0/sapphire/00011
    Then Alice will have the following tokens:
      | type | amount |
      | gold | 0      |
    And Alice will have the following reserved cards:
      | 1/0/sapphire/00011 |
    And Alice will have 0 points
    And the card in row 3 column 2 will be 1/1/diamond/10100
    And it will be Bob's turn
    And the following tokens will be available:
      | type | amount |
      | gold | 0      |

  Scenario: Alice buys a card she has in reserve
    Given Alice had the following reserved cards:
      | 3/3/emerald/07200 |
      | 3/2/emerald/07200 |
      | 3/1/emerald/07200 |
    And Alice had the following tokens:
      | type     | amount |
      | sapphire | 7      |
      | ruby     | 2      |
    When Alice buys the card 3/2/emerald/07200
    Then Alice will have 2 points
    And Alice will have the following reserved cards:
      | 3/3/emerald/07200 |
      | 3/1/emerald/07200 |
    And it will be Bob's turn
    And Alice will have the following tokens:
      | type     | amount |
      | sapphire | 0      |
      | ruby     | 0      |
