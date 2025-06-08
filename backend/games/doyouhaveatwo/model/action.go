package model

import "fmt"

type InputType string

type Input interface {
	Apply(g *Game, apply func(Event)) error
	Type() InputType
	PlayerID() *PlayerID
}

type Action interface {
	fmt.Stringer
	PlayerID() PlayerID
	IsComplete() bool
	NextActions(*Game) []Action
	ToInput() Input
	Type() string
}

type GameUpdate struct {
	Game             Game
	Event            Event
	AvailableActions map[PlayerID]Redactable[[]Action]
}
