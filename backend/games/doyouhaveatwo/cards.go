package doyouhaveatwo

import "errors"

type Card interface {
	Play(game *Game, player *Player, target *Player) error
	CanTarget(game *Game, player *Player, target *Player) bool
	Value() int
	Name() string
	Description() string
	Quantity() int
}

type Guard struct{}

func (g Guard) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (g Guard) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (g Guard) Value() int {
	return 1
}

func (g Guard) Name() string {
	return "Guard"
}

func (g Guard) Description() string {
	return "Guess another player's card. If correct, that player is eliminated."
}

func (g Guard) Quantity() int {
	return 5
}

type Priest struct{}

func (p Priest) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (p Priest) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (p Priest) Value() int {
	return 2
}

func (p Priest) Name() string {
	return "Priest"
}

func (p Priest) Description() string {
	return "Look at another player's hand."
}

func (p Priest) Quantity() int {
	return 2
}

type Baron struct{}

func (b Baron) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (b Baron) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (b Baron) Value() int {
	return 3
}

func (b Baron) Name() string {
	return "Baron"
}

func (b Baron) Description() string {
	return "Compare hands with another player. Player with lower value card is eliminated."
}

func (b Baron) Quantity() int {
	return 2
}

type Handmaid struct{}

func (h Handmaid) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (h Handmaid) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (h Handmaid) Value() int {
	return 4
}

func (h Handmaid) Name() string {
	return "Handmaid"
}

func (h Handmaid) Description() string {
	return "Player cannot be targeted by other players' cards until their next turn."
}

func (h Handmaid) Quantity() int {
	return 2
}

type Prince struct{}

func (p Prince) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (p Prince) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (p Prince) Value() int {
	return 5
}

func (p Prince) Name() string {
	return "Prince"
}

func (p Prince) Description() string {
	return "Target player discards their hand and draws a new card. Can target self."
}

func (p Prince) Quantity() int {
	return 2
}

type King struct{}

func (k King) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (k King) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (k King) Value() int {
	return 6
}

func (k King) Name() string {
	return "King"
}

func (k King) Description() string {
	return "Trade hands with another player."
}

func (k King) Quantity() int {
	return 1
}

type Countess struct{}

func (c Countess) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (c Countess) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (c Countess) Value() int {
	return 7
}

func (c Countess) Name() string {
	return "Countess"
}

func (c Countess) Description() string {
	return "No special effect, but must be discarded if holding King or Prince."
}

func (c Countess) Quantity() int {
	return 1
}

type Princess struct{}

func (p Princess) Play(game *Game, player *Player, target *Player) error {
	return errors.New("not implemented")
}

func (p Princess) CanTarget(game *Game, player *Player, target *Player) bool {
	return false
}

func (p Princess) Value() int {
	return 8
}

func (p Princess) Name() string {
	return "Princess"
}

func (p Princess) Description() string {
	return "If discarded (played or forced to discard), player is immediately eliminated."
}

func (p Princess) Quantity() int {
	return 1
}