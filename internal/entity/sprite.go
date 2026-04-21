package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 16

type Sprite struct {
	X, Y        float64
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

func (s *Sprite) Draw(screen *ebiten.Image) {
	sx := s.Frame * tileSize
	sy := s.Direction * tileSize

	frame := s.Image.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	if s.FacingRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(tileSize), 0)
	}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(frame, op)
}
