package transition

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = cols * tileSize
	screenHeight = rows * tileSize
	maxInset     = screenWidth / 2
	boxSpeed     = 8
	holdFrames   = 8
)

type BoxTransition struct {
	inset       int
	holdCounter int
	phase       Phase
	Done        bool
	hBar        *ebiten.Image // full-width horizontal bar
	vBar        *ebiten.Image // full-height vertical bar
}

func NewBox(phase Phase) *BoxTransition {
	hBar := ebiten.NewImage(screenWidth, maxInset)
	hBar.Fill(color.Black)
	vBar := ebiten.NewImage(maxInset, screenHeight)
	vBar.Fill(color.Black)
	t := &BoxTransition{
		phase: phase,
		hBar:  hBar,
		vBar:  vBar,
	}
	if phase == Opening {
		t.inset = maxInset
	}
	return t
}

func (t *BoxTransition) Update() {
	if t.Done {
		return
	}

	if t.phase == Closing {
		t.inset += boxSpeed
		if t.inset >= maxInset {
			t.inset = maxInset
			t.holdCounter++
			if t.holdCounter >= holdFrames {
				t.Done = true
			}
		}
	} else {
		t.inset -= boxSpeed
		if t.inset <= 0 {
			t.inset = 0
			t.Done = true
		}
	}
}

func (t *BoxTransition) IsFullyBlack() bool {
	if t.phase == Opening {
		return true
	}
	return t.inset >= maxInset
}

func (t *BoxTransition) IsDone() bool { return t.Done }

func (t *BoxTransition) Draw(screen *ebiten.Image) {
	if t.inset <= 0 {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// top bar
	op.GeoM.Reset()
	op.GeoM.Scale(1, float64(t.inset)/float64(maxInset))
	screen.DrawImage(t.hBar, op)

	// bottom bar
	op.GeoM.Reset()
	op.GeoM.Scale(1, float64(t.inset)/float64(maxInset))
	op.GeoM.Translate(0, float64(screenHeight-t.inset))
	screen.DrawImage(t.hBar, op)

	// left bar
	op.GeoM.Reset()
	op.GeoM.Scale(float64(t.inset)/float64(maxInset), 1)
	op.GeoM.Translate(0, float64(t.inset))
	screen.DrawImage(t.vBar, op)

	// right bar
	op.GeoM.Reset()
	op.GeoM.Scale(float64(t.inset)/float64(maxInset), 1)
	op.GeoM.Translate(float64(screenWidth-t.inset), float64(t.inset))
	screen.DrawImage(t.vBar, op)
}
