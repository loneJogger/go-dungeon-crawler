package entity

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
)

const speed = 1.5

type Player struct {
	Sprite
	dirStack []ebiten.Key
}

func NewPlayer(x, y float64, image *ebiten.Image) *Player {
	return &Player{Sprite: Sprite{X: x, Y: y, Image: image}}
}

func (p *Player) Update(m *tiled.Map, npcs []*NPC) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// push newly pressed direction keys onto stack
	for _, k := range []ebiten.Key{ebiten.KeyArrowDown, ebiten.KeyArrowUp, ebiten.KeyArrowLeft, ebiten.KeyArrowRight} {
		if inpututil.IsKeyJustPressed(k) {
			p.dirStack = append(p.dirStack, k)
		}
	}
	// remove released keys from stack
	held := p.dirStack[:0]
	for _, k := range p.dirStack {
		if ebiten.IsKeyPressed(k) {
			held = append(held, k)
		}
	}
	p.dirStack = held

	nx, ny := p.X, p.Y
	p.Moving = len(p.dirStack) > 0

	if p.Moving {
		switch p.dirStack[len(p.dirStack)-1] {
		case ebiten.KeyArrowDown:
			ny += speed
			p.Direction = 0
			p.FacingRight = false
		case ebiten.KeyArrowUp:
			ny -= speed
			p.Direction = 2
			p.FacingRight = false
		case ebiten.KeyArrowLeft:
			nx -= speed
			p.Direction = 1
			p.FacingRight = false
		case ebiten.KeyArrowRight:
			nx += speed
			p.Direction = 1
			p.FacingRight = true
		}

		if !p.OutOfBounds(m, nx, ny) && !p.IsSolidAt(m, nx, ny) && !npcCollides(nx, ny, npcs) {
			p.X, p.Y = nx, ny
		}
	}

	p.TickAnim()
}


func npcCollides(x, y float64, npcs []*NPC) bool {
	bounds := image.Rect(int(x), int(y), int(x)+tileSize, int(y)+tileSize)
	for _, npc := range npcs {
		if bounds.Overlaps(npc.Bounds()) {
			return true
		}
	}
	return false
}
