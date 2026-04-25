package transition

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Phase int

const (
	Closing Phase = iota
	Opening
)

type TransitionType int

const (
	TransitionSpiral TransitionType = iota
	TransitionBox
)

type Transition interface {
	Update()
	Draw(screen *ebiten.Image)
	IsFullyBlack() bool
	IsDone() bool
}

const (
	cols         = 20
	rows         = 15
	tileSize     = 16
	speed        = 1
	stepsPerTick = 3
)

type SpiralTransition struct {
	tiles []bool
	order []int
	step  int
	timer int
	phase Phase
	Done  bool
	black *ebiten.Image
}

func New(phase Phase) *SpiralTransition {
	t := &SpiralTransition{
		tiles: make([]bool, cols*rows),
		phase: phase,
	}
	t.black = ebiten.NewImage(tileSize, tileSize)
	t.black.Fill(color.Black)
	t.order = generateSpiralOrder()
	if phase == Opening {
		for i := range t.tiles {
			t.tiles[i] = true
		}
	}
	return t
}

func (t *SpiralTransition) Update() {
	if t.Done {
		return
	}

	t.timer++
	if t.timer < speed {
		return
	}
	t.timer = 0

	if t.phase == Closing {
		for range stepsPerTick {
			if t.step < len(t.order) {
				t.tiles[t.order[t.step]] = true
				t.step++
			} else {
				// fully black, switch to opening
				t.phase = Opening
				t.step = 0
			}
		}
	} else {
		for range stepsPerTick {
			if t.step < len(t.order) {
				t.tiles[t.order[t.step]] = false
				t.step++
			} else {
				t.Done = true
			}
		}
	}
}

func (t *SpiralTransition) Draw(screen *ebiten.Image) {

	for i, solid := range t.tiles {
		if !solid {
			continue
		}
		x := (i % cols) * tileSize
		y := (i / cols) * tileSize
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(t.black, op)
	}
}

func (t *SpiralTransition) IsFullyBlack() bool {
	return t.phase == Opening
}

func (t *SpiralTransition) IsDone() bool { return t.Done }

func generateSpiralOrder() []int {
	order := make([]int, 0, cols*rows)
	added := make(map[int]bool)
	cx, cy := cols/2, rows/2
	add := func(x, y int) {
		if x < 0 || x >= cols || y < 0 || y >= rows {
			return
		}
		idx := y*cols + x
		if !added[idx] {
			added[idx] = true
			order = append(order, idx)
		}
	}
	add(cx-1, cy)
	add(cx, cy)
	add(cx-1, cy-1)
	add(cx, cy-1)
	for r := 1; r < max(cols, rows); r++ {
		// top
		for x := cx - r; x <= cx+r; x++ {
			add(x, cy-r)
		}
		// right
		for y := cy - r + 1; y <= cy+r; y++ {
			add(cx+r, y)
		}
		// bottom
		for x := cx + r - 1; x >= cx-r; x-- {
			add(x, cy+r)
		}
		// left
		for y := cy + r - 1; y >= cy-r+1; y-- {
			add(cx-r, y)
		}
	}
	return order
}
