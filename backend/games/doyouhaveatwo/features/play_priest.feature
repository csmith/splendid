Feature: Play Priest

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Priest |
      | Baron  |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Baron |
    And the game has phase "play"

  Scenario: Player is able to perform action play_priest when it's their turn and they have a Priest
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest", "targetPlayer": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And the following event occurred: "hand_revealed"

  Scenario: Player cannot perform action play_priest when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    Then an error occurs

  Scenario: Player cannot perform action play_priest when they don't have a Priest
    Given it is player B's turn
    When player B sends action {"type": "play_card_target_others", "player": "B", "cardName": "Priest"}
    Then an error occurs

  Scenario: Cannot target protected player with Priest
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target eliminated player with Priest
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target yourself with Priest
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest", "targetPlayer": "A"}
    Then an error occurs

  Scenario: Available actions when player holds Priest
    Given it is player A's turn
    Then the available actions should be:
      | player | action                |
      | A      | play_priest(player=A) |
      | A      | play_baron(player=A)  |

  Scenario: Available actions when player starts playing Priest
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_priest(player=A, target=B) |
      | A      | play_priest(player=A, target=C) |

  Scenario: Available actions exclude protected players as targets
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_priest(player=A, target=C) |

  Scenario: Available actions exclude eliminated players as targets
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "Priest"}
    Then no error occurs
    And the available actions should be:
      | player | action                          |
      | A      | play_priest(player=A, target=C) |

  Scenario: Available actions offer discard when no valid targets exist
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    Then the available actions should be:
      | player | action                              |
      | A      | discard_card(player=A, card=Priest) |
      | A      | discard_card(player=A, card=Baron) |

  Scenario: Player can discard Priest when no valid targets
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    When player A sends action {"type": "discard_card", "player": "A", "cardName": "Priest"}
    Then no error occurs
    And player A should have 1 cards in discard pile