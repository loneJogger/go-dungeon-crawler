package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const tileSize = 16
const speed = 2.0

type Player struct {
	X, Y       float64
	Direction  int
	Frame      int
	FacingLeft bool
	Moving     bool
	AnimTimer  int
}

func NewPlayer(x, y float64) *Player {
	return &Player{X: x, Y: y}
}

func (p *Player) Update(m *tiled.Map) {
	nx, ny := p.X, p.Y

	// walk down
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ny += speed
		p.Direction = 0
		p.FacingLeft = false
	}

	// walk up
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		ny -= speed
		p.Direction = 2
		p.FacingLeft = false
	}

	// walk left
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		nx -= speed
		p.Direction = 1
		p.FacingLeft = true
	}

	// walk right
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		nx += speed
		p.Direction = 1
		p.FacingLeft = false
	}

	// is moving check
	p.Moving = ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowRight)

	// check for collision
	if !world.IsSolid(m, nx+8, ny+8) {
		p.X, p.Y = nx, ny
	}

	// animation timer
	if p.Moving {
		p.AnimTimer++
		if p.AnimTimer >= 10 {
			p.AnimTimer = 0
			if p.Frame == 0 {
				p.Frame = 1
			} else {
				p.Frame = 0
			}
		}
	} else {
		p.Frame = 0
		p.AnimTimer = 0
	}
}

func (p *Player) Draw(screen *ebiten.Image, sprite *ebiten.Image) {
	sx := p.Frame * tileSize
	sy := p.Direction * tileSize

	frame := sprite.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}

	if p.FacingLeft {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(tileSize, 0)
	}

	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(frame, op)
}
