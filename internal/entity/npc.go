package entity

type NPC struct {
	Sprite
	OnInteract func()
}

func NewNPC(x, y float64, direction int) *NPC {
	return &NPC{
		Sprite: Sprite{X: x, Y: y, Direction: direction},
	}
}
