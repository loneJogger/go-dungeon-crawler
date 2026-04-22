package entity

import (
	"math/rand"

	"github.com/lafriks/go-tiled"
)

const npcSpeed = 0.6
const wanderInterval = 90

type NPC struct {
	Sprite
	OnInteract  func()
	Wanders     bool
	wanderTimer int
	wanderDir   int
	paused      bool
}

func NewNPC(x, y float64, direction int) *NPC {
	return &NPC{
		Sprite: Sprite{X: x, Y: y, Direction: direction},
	}
}

func (n *NPC) Update(m *tiled.Map) {
	if !n.Wanders {
		return
	}

	n.wanderTimer++
	if n.wanderTimer >= wanderInterval {
		n.wanderTimer = 0
		choice := rand.Intn(5) // 0-3 = directions, 4 = pause
		n.paused = choice == 4
		n.wanderDir = choice
	}

	if n.paused {
		n.Moving = false
		return
	}

	var dx, dy float64
	switch n.wanderDir {
	case 0:
		dy = npcSpeed
		n.Direction = 0
		n.FacingRight = false
	case 1:
		dx = -npcSpeed
		n.Direction = 1
		n.FacingRight = false
	case 2:
		dy = -npcSpeed
		n.Direction = 2
		n.FacingRight = false
	case 3:
		dx = npcSpeed
		n.Direction = 1
		n.FacingRight = true
	}

	nx, ny := n.X+dx, n.Y+dy
	if !n.OutOfBounds(m, nx, ny) && !n.IsSolidAt(m, nx, ny) {
		n.X, n.Y = nx, ny
		n.Moving = true
	} else {
		n.Moving = false
		n.wanderTimer = wanderInterval // force a new direction next frame
	}

	n.TickAnim()
}
