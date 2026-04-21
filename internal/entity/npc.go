package entity

type InteractionType int

const (
	InteractionDialog InteractionType = iota
	InteractionBattle
)

type NPC struct {
	Sprite
	Interaction InteractionType
}

func NewNPC(x, y float64, direction int, interaction InteractionType) *NPC {
	return &NPC{
		Sprite:      Sprite{X: x, Y: y, Direction: direction},
		Interaction: interaction,
	}
}
