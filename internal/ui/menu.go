package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func playSound(p *audio.Player) {
	if p == nil {
		return
	}
	p.Rewind()
	p.Play()
}

func (m *Menu) Draw(screen *ebiten.Image, x, y int) {
	for i, item := range m.Items {
		label := "  " + item.Label
		if i == m.focused {
			label = "> " + item.Label
		}
		ebitenutil.DebugPrintAt(screen, label, x, y+i*DialogLineHeight)
	}
}
