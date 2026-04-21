package entity

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const speed = 1.5

type Player struct {
	Sprite
}

func NewPlayer(x, y float64, image *ebiten.Image) *Player {
	return &Player{Sprite: Sprite{X: x, Y: y, Image: image}}
}

func (p *Player) Update(m *tiled.Map) {
	nx, ny := p.X, p.Y

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ny += speed
		p.Direction = 0
		p.FacingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		ny -= speed
		p.Direction = 2
		p.FacingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		nx -= speed
		p.Direction = 1
		p.FacingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		nx += speed
		p.Direction = 1
		p.FacingRight = true
	}

	p.Moving = ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowRight)

	var checkX, checkY float64 = nx + 8, ny + 8
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		checkY = ny + tileSize - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		checkY = ny + 8
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		checkX = nx
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		checkX = nx + tileSize - 1
	}

	if !world.IsSolid(m, checkX, checkY) {
		p.X, p.Y = nx, ny
	}

	p.TickAnim()
}
