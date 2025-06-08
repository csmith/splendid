package model

type Card interface {
	Play(game *Game, player *Player, target *Player) error
	CanTarget(game *Game, player *Player, target *Player) bool
	Value() int
	Name() string
	Description() string
	Quantity() int
}
