package model

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	value       int
	name        string
	description string
	quantity    int
}

func (c Card) Value() int {
	return c.value
}

func (c Card) Name() string {
	return c.name
}

func (c Card) Description() string {
	return c.description
}

func (c Card) Quantity() int {
	return c.quantity
}

func (c Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.name)
}

func (c Card) String() string {
	return c.name
}

var (
	CardGuard    = Card{value: 1, name: "Guard", description: "Guess another player's card. If correct, that player is eliminated.", quantity: 5}
	CardPriest   = Card{value: 2, name: "Priest", description: "Look at another player's hand.", quantity: 2}
	CardBaron    = Card{value: 3, name: "Baron", description: "Compare hands with another player. Player with lower value card is eliminated.", quantity: 2}
	CardHandmaid = Card{value: 4, name: "Handmaid", description: "Player cannot be targeted by other players' cards until their next turn.", quantity: 2}
	CardPrince   = Card{value: 5, name: "Prince", description: "Target player discards their hand and draws a new card. Can target self.", quantity: 2}
	CardKing     = Card{value: 6, name: "King", description: "Trade hands with another player.", quantity: 1}
	CardCountess = Card{value: 7, name: "Countess", description: "No special effect, but must be discarded if holding King or Prince.", quantity: 1}
	CardPrincess = Card{value: 8, name: "Princess", description: "If discarded (played or forced to discard), player is immediately eliminated.", quantity: 1}
)

var CardTypes = []Card{
	CardGuard, CardPriest, CardBaron, CardHandmaid,
	CardPrince, CardKing, CardCountess, CardPrincess,
}

// GetCardByName returns the card with the given name, or an error if not found
func GetCardByName(name string) (Card, error) {
	for _, card := range CardTypes {
		if card.Name() == name {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("unknown card name: %s", name)
}
