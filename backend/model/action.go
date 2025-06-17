package model

import (
	"fmt"
)

// Action represents a multi-step player interaction that may require several rounds
// of client-server communication to complete. It is generic over the game type G
// and input type I, allowing it to be used by different game implementations.
type Action[G any, I any] interface {
	fmt.Stringer
	Typeable
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(G) []Action[G, I]
	ToInput() I
}
