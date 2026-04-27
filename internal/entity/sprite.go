package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/loneJogger/go-dungeon-crawler/internal/scene/explore"
)

const tileSize = 16

const (
	collideLeft   = 2
	collideRight  = 13
	collideTop    = 10
	collideBottom = 15
)

type Sprite struct {
	X, Y        float64
	OffsetX     float64
	OffsetY     float64
	Direction   int
	Frame       int
	FacingRight bool
	Moving      bool
	AnimTimer   int
	Image       *ebiten.Image
}

func (s *Sprite) Bounds() image.Rectangle {
	return image.Rect(int(s.X), int(s.Y), int(s.X)+tileSize, int(s.Y)+tileSize)
}

func (s *Sprite) TickAnim() {
	if s.Moving {
		s.AnimTimer++
		if s.AnimTimer >= 10 {
			s.AnimTimer = 0
			s.Frame ^= 1
		}
	} else {
		s.Frame = 0
		s.AnimTimer = 0
	}
}

func (s *Sprite) OutOfBounds(m *tiled.Map, nx, ny float64) bool {
	mapW := float64(m.Width * m.TileWidth)
	mapH := float64(m.Height * m.TileHeight)
	return nx < 0 || nx > mapW-tileSize || ny < 0 || ny > mapH-tileSize
}

func (s *Sprite) IsSolidAt(m *tiled.Map, nx, ny float64) bool {
	switch s.Direction {
	case 0: // down
		return explore.IsCollison(m, nx+collideLeft, ny+collideBottom) ||
			explore.IsCollison(m, nx+collideRight, ny+collideBottom)
	case 2: // up
		return explore.IsCollison(m, nx+collideLeft, ny+collideTop) ||
			explore.IsCollison(m, nx+collideRight, ny+collideTop)
	case 1:
		if !s.FacingRight {
			return explore.IsCollison(m, nx+collideLeft, ny+collideTop) ||
				explore.IsCollison(m, nx+collideLeft, ny+collideBottom)
		}
		return explore.IsCollison(m, nx+collideRight, ny+collideTop) ||
			explore.IsCollison(m, nx+collideRight, ny+collideBottom)
	}
	return false
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	sx := s.Frame * tileSize
	sy := s.Direction * tileSize

	frame := s.Image.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	if s.FacingRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(tileSize), 0)
	}
	op.GeoM.Translate(s.X-s.OffsetX, s.Y-s.OffsetY)
	screen.DrawImage(frame, op)
}
