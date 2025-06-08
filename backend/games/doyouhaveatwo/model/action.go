package model

type Action struct {
	Type  string
	Value interface{}
	Label string
}

type GameUpdate struct {
	Game             Game
	Event            Event
	AvailableActions map[PlayerID]Redactable[[]Action]
}
