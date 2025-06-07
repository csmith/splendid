package doyouhaveatwo

type EventType string

const (
	EventStartRound EventType = "start_round"
)

type Event struct {
	Type     EventType
	PlayerID *PlayerID
	Details  interface{}
}