package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuStack struct {
	stack       []*Menu
	cancelSound *audio.Player
}

func NewMenuStack(root *Menu, sounds MenuSounds) *MenuStack {
	root.NavSound = sounds.Nav
	root.SelectSound = sounds.Select
	return &MenuStack{stack: []*Menu{root}, cancelSound: sounds.Cancel}
}

func (ms *MenuStack) Active() *Menu {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MenuStack) Stack() []*Menu {
	return ms.stack
}

func (ms *MenuStack) Push(menu *Menu) {
	ms.stack = append(ms.stack, menu)
}

func (ms *MenuStack) Pop() {
	if len(ms.stack) > 1 {
		ms.stack = ms.stack[:len(ms.stack)-1]
		playSound(ms.cancelSound)
	}
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
