Feature: Play King

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | King  |
      | Guard |
    And player B has the following cards in their hand:
      | Guard |
    And player C has the following cards in their hand:
      | Priest |
    And the game has phase "play"

  Scenario: Player is able to perform action play_king when it's their turn and they have a King
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King", "targetPlayer": "B"}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And the following event occurred: "hands_swapped"

  Scenario: King trades hands correctly
    Given it is player A's turn
    And player A has the following cards in their hand:
      | King  |
      | Baron |
    And player B has the following cards in their hand:
      | Princess |
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King", "targetPlayer": "B"}
    Then no error occurs
    And player A should have card "Princess" in their hand
    And player B should have card "Baron" in their hand

  Scenario: Player cannot perform action play_king when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    Then an error occurs

  Scenario: Player cannot perform action play_king when they don't have a King
    Given it is player B's turn
    When player B sends action {"type": "play_card_target_others", "player": "B", "cardName": "King"}
    Then an error occurs

  Scenario: Cannot target protected player with King
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target eliminated player with King
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King", "targetPlayer": "B"}
    Then an error occurs

  Scenario: Cannot target yourself with King
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    And player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King", "targetPlayer": "A"}
    Then an error occurs

  Scenario: Available actions when player holds King
    Given it is player A's turn
    Then the available actions should be:
      | player | action               |
      | A      | play_king(player=A)  |
      | A      | play_guard(player=A) |

  Scenario: Available actions when player starts playing King
    Given it is player A's turn
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    Then no error occurs
    And the available actions should be:
      | player | action                        |
      | A      | play_king(player=A, target=B) |
      | A      | play_king(player=A, target=C) |

  Scenario: Available actions exclude protected players as targets
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    Then no error occurs
    And the available actions should be:
      | player | action                        |
      | A      | play_king(player=A, target=C) |

  Scenario: Available actions exclude eliminated players as targets
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "play_card_target_others", "player": "A", "cardName": "King"}
    Then no error occurs
    And the available actions should be:
      | player | action                        |
      | A      | play_king(player=A, target=C) |

  Scenario: Available actions offer discard when no valid targets exist
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    Then the available actions should be:
      | player | action                             |
      | A      | discard_card(player=A, card=King)  |
      | A      | discard_card(player=A, card=Guard) |

  Scenario: Player can discard King when no valid targets
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    When player A sends action {"type": "discard_card", "player": "A", "cardName": "King"}
    Then no error occurs
    And player A should have 1 cards in discard pile