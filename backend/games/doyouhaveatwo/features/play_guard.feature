Feature: Play Guard

  Background:
    Given a game with 3 players
    And a round starts
    And player A has the following cards in their hand:
      | Guard |
      | King  |
    And player B has the following cards in their hand:
      | Priest |
    And player C has the following cards in their hand:
      | Baron |
    And the game has phase "play"

  Scenario: Player is able to perform action play_guard when it's their turn and they have a Guard
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B", "guessed_rank": 2}
    Then no error occurs
    And player A should have 1 cards in discard pile
    And player B should be eliminated
    And player B should have 1 cards in discard pile

  Scenario: Player cannot perform action play_guard when it's not their turn
    Given it is player B's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    Then an error occurs

  Scenario: Player cannot perform action play_guard when they don't have a Guard
    Given it is player B's turn
    When player B sends action {"type": "dyhat:a:play_card_with_guess", "player": "B", "card_name": "Guard"}
    Then an error occurs

  Scenario: Play Guard with incorrect guess does not eliminate target
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B", "guessed_rank": 3}
    Then no error occurs
    And player B should not be eliminated
    And player A should have 1 cards in discard pile

  Scenario: Play Guard with correct guess eliminates target
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B", "guessed_rank": 2}
    Then no error occurs
    And player B should be eliminated
    And player A should have 1 cards in discard pile

  Scenario: Cannot target protected player with Guard
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    Then an error occurs

  Scenario: Cannot target eliminated player with Guard
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    Then an error occurs

  Scenario: Cannot target yourself with Guard
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "A"}
    Then an error occurs

  Scenario: Cannot guess Guard with Guard
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B", "guessed_rank": 1}
    Then an error occurs

  Scenario: Available actions when player holds Guard
    Given it is player A's turn
    Then the available actions should be:
      | player | action               |
      | A      | play_guard(player=A) |
      | A      | play_king(player=A)  |

  Scenario: Available actions when player starts playing Guard
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_guard(player=A, target=B) |
      | A      | play_guard(player=A, target=C) |

  Scenario: Available actions when player selects Guard target
    Given it is player A's turn
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    And player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard", "target_player": "B"}
    Then no error occurs
    And the available actions should be:
      | player | action                                  |
      | A      | play_guard(player=A, target=B, guess=2) |
      | A      | play_guard(player=A, target=B, guess=3) |
      | A      | play_guard(player=A, target=B, guess=4) |
      | A      | play_guard(player=A, target=B, guess=5) |
      | A      | play_guard(player=A, target=B, guess=6) |
      | A      | play_guard(player=A, target=B, guess=7) |
      | A      | play_guard(player=A, target=B, guess=8) |

  Scenario: Available actions exclude protected players as targets
    Given it is player A's turn
    And player B is protected
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_guard(player=A, target=C) |

  Scenario: Available actions exclude eliminated players as targets
    Given it is player A's turn
    And player B is eliminated
    When player A sends action {"type": "dyhat:a:play_card_with_guess", "player": "A", "card_name": "Guard"}
    Then no error occurs
    And the available actions should be:
      | player | action                         |
      | A      | play_guard(player=A, target=C) |

  Scenario: Available actions offer discard when no valid targets exist
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    Then the available actions should be:
      | player | action                             |
      | A      | discard_card(player=A, card=Guard) |
      | A      | discard_card(player=A, card=King)  |

  Scenario: Player can discard Guard when no valid targets
    Given it is player A's turn
    And player B is eliminated
    And player C is protected
    When player A sends action {"type": "dyhat:a:discard_card", "player": "A", "card_name": "Guard"}
    Then no error occurs
    And player A should have 1 cards in discard pile