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

type MenuStack struct {
	stack       []*Menu
	CancelSound *audio.Player
}

func NewMenuStack(root *Menu) *MenuStack {
	return &MenuStack{stack: []*Menu{root}}
}

func (ms *MenuStack) Push(menu *Menu) {
	ms.stack = append(ms.stack, menu)
}

func (ms *MenuStack) Pop() {
	if len(ms.stack) > 1 {
		ms.stack = ms.stack[:len(ms.stack)-1]
		playSound(ms.CancelSound)
	}
}

func (ms *MenuStack) Active() *Menu {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MenuStack) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		ms.Pop()
		return
	}
	ms.Active().Update()
}

func (ms *MenuStack) Draw(screen *ebiten.Image, font *ebiten.Image, x, y int) {
	ms.Active().Draw(screen, font, x, y)
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
