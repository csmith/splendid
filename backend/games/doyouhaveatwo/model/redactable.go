package model

type Redactable[T any] struct {
	Value     T
	VisibleTo map[PlayerID]bool
}