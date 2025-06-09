Feature: Play Prince

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Prince |
      | Baron  |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Player is able to perform action play_prince when it's their turn and they have a Prince
    Given it is player A's turn
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should have 1 cards in discard pile
    And the following event occurred: "card_dealt"

  Scenario: Prince can target self
    Given it is player A's turn
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "A"}
    Then no error occurs
    And player A should have 2 cards in discard pile
    And the following event occurred: "card_dealt"

  Scenario: Prince eliminates target if they discard Princess
    Given it is player A's turn
    And player B has the following cards in their hand:
      | Princess |
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should be eliminated
    And player B should have 1 cards in discard pile

  Scenario: Prince forces player to draw removed card when deck is empty
    Given it is player A's turn
    And the deck is empty
    And the removed card is a Princess
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "B"}
    Then no error occurs
    And player B should have card "Princess" in their hand

  Scenario: Player cannot perform action play_prince when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then an error occurs

  Scenario: Player cannot perform action play_prince when they don't have a Prince
    Given it is player B's turn
    When player B sends action {"type": "play_prince", "player": "B", "card_name": "Prince"}
    Then an error occurs

  Scenario: Cannot target protected player with Prince
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "B"}
    Then an error occurs

  Scenario: Cannot target eliminated player with Prince
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "B"}
    Then an error occurs

  Scenario: Can target yourself with Prince
    Given it is player A's turn
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    And player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince", "target_player": "A"}
    Then no error occurs

  Scenario: Must target self when all other players are protected
    Given it is player A's turn
    And player B is protected
    And player C is protected
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then the available actions should be:
      | player | action                          |
      | A      | play_prince(player=A, target=A) |

  Scenario: Available actions when player holds Prince
    Given it is player A's turn
    Then the available actions should be:
      | player | action                |
      | A      | play_prince(player=A) |
      | A      | play_baron(player=A)  |

  Scenario: Available actions when player starts playing Prince
    Given it is player A's turn
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_prince(player=A, target=A) |
      | A      | play_prince(player=A, target=B) |
      | A      | play_prince(player=A, target=C) |

  Scenario: Available actions exclude protected players as targets
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_prince(player=A, target=A) |
      | A      | play_prince(player=A, target=C) |

  Scenario: Available actions exclude eliminated players as targets
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_prince", "player": "A", "card_name": "Prince"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_prince(player=A, target=A) |
      | A      | play_prince(player=A, target=C) |