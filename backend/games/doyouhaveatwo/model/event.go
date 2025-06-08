package model

type EventType string

type Event interface {
	Type() EventType
	PlayerID() *PlayerID
	Apply(g *Game) error
}