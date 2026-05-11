package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	Items       []MenuItem
	focused     int
	blinkTick   int
	Static      bool
	ItemGap     int
	NavSound    *audio.Player
	SelectSound *audio.Player
}

type MenuItem struct {
	Label    string
	OnSelect func()
	HeightPx int // 0 falls back to DialogLineHeight
	DrawItem func(screen *ebiten.Image, font *ebiten.Image, x, y int, focused, cursorOn bool)
}

type MenuSounds struct {
	Nav    *audio.Player
	Select *audio.Player
	Cancel *audio.Player
}

func NewMenu(items []MenuItem) *Menu {
	return &Menu{Items: items}
}

func (m *Menu) MoveUp() {
	m.focused = (m.focused - 1 + len(m.Items)) % len(m.Items)
	m.blinkTick = 0
}

func (m *Menu) MoveDown() {
	m.focused = (m.focused + 1) % len(m.Items)
	m.blinkTick = 0
}

func (m *Menu) Select() { m.Items[m.focused].OnSelect() }

func (m *Menu) Focused() int { return m.focused }

func playSound(p *audio.Player) {
	if p == nil {
		return
	}
	p.Rewind()
	p.Play()
}

func (m *Menu) Update() {
	m.blinkTick++
	if !m.Static {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			m.MoveUp()
			playSound(m.NavSound)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			m.MoveDown()
			playSound(m.NavSound)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		m.Select()
		playSound(m.SelectSound)
	}
}

func (m *Menu) Draw(screen *ebiten.Image, font *ebiten.Image, x, y int) {
	cursorOn := (m.blinkTick/30)%2 == 0
	rowY := y
	for i, item := range m.Items {
		if item.DrawItem != nil {
			item.DrawItem(screen, font, x, rowY, i == m.focused, cursorOn)
		} else {
			cursor := " "
			if i == m.focused && cursorOn {
				cursor = ">"
			}
			DrawText(screen, font, cursor+item.Label, x, rowY)
		}
		h := DialogLineHeight
		if item.HeightPx > 0 {
			h = item.HeightPx
		}
		rowY += h + m.ItemGap
	}
}

// DrawText renders a plain ASCII string at (x, y) using the sprite font.
func DrawText(screen *ebiten.Image, font *ebiten.Image, text string, x, y int) {
	for col, ch := range text {
		charIndex := int(ch) - 32
		sx := (charIndex % charsPerRow) * charSize
		sy := (charIndex / charsPerRow) * charSize
		src := font.SubImage(image.Rect(sx, sy, sx+charSize, sy+charSize)).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x+col*charSize), float64(y))
		screen.DrawImage(src, op)
	}
}
