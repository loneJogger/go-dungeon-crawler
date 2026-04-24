package transition

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const wipeSpeed = 16

type WipeDirection int

const (
	WipeUp WipeDirection = iota
	WipeDown
)

type WipeTransition struct {
	curtainH  int
	direction WipeDirection
	done      bool
	overlay   *ebiten.Image
}

func NewWipe(direction WipeDirection) *WipeTransition {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.Black)
	t := &WipeTransition{direction: direction, overlay: overlay}
	if direction == WipeDown {
		t.curtainH = screenHeight
	}
	return t
}

func (t *WipeTransition) Update() {
	if t.done {
		return
	}
	if t.direction == WipeUp {
		t.curtainH += wipeSpeed
		if t.curtainH >= screenHeight {
			t.curtainH = screenHeight
			t.done = true
		}
	} else {
		t.curtainH -= wipeSpeed
		if t.curtainH <= 0 {
			t.curtainH = 0
			t.done = true
		}
	}
}

func (t *WipeTransition) Draw(screen *ebiten.Image) {
	if t.curtainH <= 0 {
		return
	}
	src := t.overlay.SubImage(t.overlay.Bounds()).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}
	scaleY := float64(t.curtainH) / float64(screenHeight)
	op.GeoM.Scale(1, scaleY)
	op.GeoM.Translate(0, float64(screenHeight-t.curtainH))
	screen.DrawImage(src, op)
}

func (t *WipeTransition) IsFullyBlack() bool { return t.curtainH >= screenHeight }
func (t *WipeTransition) IsDone() bool        { return t.done }
