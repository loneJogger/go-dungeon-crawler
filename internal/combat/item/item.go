package item

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Item struct {
	Name        string
	Description string
	Icon        *ebiten.Image
	Value       int
}

type Equipement struct {
	Item
}
