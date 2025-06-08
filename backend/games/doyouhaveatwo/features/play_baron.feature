Feature: Play Baron

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Baron |
      | King  |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Player is able to perform action play_baron when it's their turn and they have a Baron
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should be eliminated

  Scenario: Baron with tie does not eliminate anyone
    Given it is player A's turn
    And player B has the following cards in their hand:
      | King |
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should not be eliminated
    And player A should not be eliminated

  Scenario: Baron eliminates player with lower value card
    Given it is player A's turn
    And player B has the following cards in their hand:
      | Princess |
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "B"}
    Then no error occurs
    And player A should have 2 cards in discard pile
    And player A should be eliminated
    And player B should not be eliminated

  Scenario: Player cannot perform action play_baron when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    Then an error occurs

  Scenario: Player cannot perform action play_baron when they don't have a Baron
    Given it is player B's turn
    When player B sends action {"type": "play_card_target_others", "player": "B", "cardName": "Baron"}
    Then an error occurs

  Scenario: Cannot target protected player with Baron
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target eliminated player with Baron
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target yourself with Baron
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron", "targetPlayer": "A"}
    Then an error occurs

  Scenario: Available actions when player holds Baron
    Given it is player A's turn
    Then the available actions should be:
      | player | action               |
      | A      | play_baron(player=A) |
      | A      | play_king(player=A)  |

  Scenario: Available actions when player starts playing Baron
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_baron(player=A, target=B) |
      | A      | play_baron(player=A, target=C) |

  Scenario: Available actions exclude protected players as targets
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_baron(player=A, target=C) |

  Scenario: Available actions exclude eliminated players as targets
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Baron"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_baron(player=A, target=C) |

  Scenario: Available actions offer discard when no valid targets exist
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    Then the available actions should be:
      | player | action                             |
      | A      | discard_card(player=A, card=Baron) |
      | A      | discard_card(player=A, card=King) |

  Scenario: Player can discard Baron when no valid targets
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    When player A sends action {"type": "discard_card", "player": "A", "cardName": "Baron"}
    Then no error occurs
    And player A should have 1 cards in discard pile