package combat

import "github.com/hajimehoshi/ebiten/v2"

type ItemType int

const (
	Consumable ItemType = iota
	Equipment
	KeyItem
)

type Item struct {
	Name        string
	Type        ItemType
	Description string
	Icon        *ebiten.Image
	Value       int
	OnUse       func(*Character)
	OnlyBattle  bool
}

type Equipement struct {
	Item
}
