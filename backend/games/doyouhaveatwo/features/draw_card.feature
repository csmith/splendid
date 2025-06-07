Feature: Draw Card
  Background:
    Given a game with 3 players
    And the game has phase "play"
    And a round starts

  Scenario: Player draws a card from deck
    When player A draws a card
    Then player A should have exactly 2 card in their hand
    And the deck should have 11 cards remaining
    And player A's cards should only be visible to themselves

  Scenario: Card becomes visible to drawing player
    When player B draws a card
    Then player B's cards should only be visible to themselves

  Scenario: Multiple players can draw cards
    When player A draws a card
    And player B draws a card
    And player C draws a card
    Then player A should have exactly 2 card in their hand
    And player B should have exactly 2 card in their hand
    And player C should have exactly 2 card in their hand
    And the deck should have 9 cards remaining

  Scenario: Cannot draw from empty deck
    Given the deck is empty
    When player A draws a card
    Then an error is returned: "failed to apply event draw_card: deck is empty, cannot draw card"