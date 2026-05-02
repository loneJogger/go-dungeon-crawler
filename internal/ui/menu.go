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
	NavSound    *audio.Player
	SelectSound *audio.Player
}

type MenuItem struct {
	Label    string
	OnSelect func()
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
}

func (m *Menu) MoveDown() {
	m.focused = (m.focused + 1) % len(m.Items)
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
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.MoveUp()
		playSound(m.NavSound)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.MoveDown()
		playSound(m.NavSound)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		m.Select()
		playSound(m.SelectSound)
	}
}

func (m *Menu) Draw(screen *ebiten.Image, font *ebiten.Image, x, y int) {
	for i, item := range m.Items {
		label := "  " + item.Label
		if i == m.focused {
			label = "> " + item.Label
		}
		col := 0
		for _, ch := range label {
			charIndex := int(ch) - 32
			sx := (charIndex % charsPerRow) * charSize
			sy := (charIndex / charsPerRow) * charSize
			src := font.SubImage(image.Rect(sx, sy, sx+charSize, sy+charSize)).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x+col*charSize), float64(y+i*DialogLineHeight))
			screen.DrawImage(src, op)
			col++
		}
	}
}
