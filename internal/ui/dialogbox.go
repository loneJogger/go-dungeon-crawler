package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DialogBoxX      = 0
	DialogBoxY      = 144
	DialogBoxWidth  = 320
	DialogBoxHeight = 96
)

type DialogMode int

const (
	DialogModeText DialogMode = iota
	DialogModeMenu
)

type DialogBox struct {
	mode       DialogMode
	textScroll *TextScroll
	menu       *Menu
	font       *ebiten.Image
	border     *ebiten.Image
	bg         *ebiten.Image
	Active     bool
}

func NewDialogBox(font *ebiten.Image, border *ebiten.Image) *DialogBox {
	bg := ebiten.NewImage(DialogBoxWidth, DialogBoxHeight)
	bg.Fill(color.Black)
	return &DialogBox{font: font, border: border, bg: bg}
}

func (d *DialogBox) ShowText(text string, onComplete func()) {
	d.textScroll = NewTextScroll(text, onComplete)
	d.mode = DialogModeText
	d.Active = true
}

func (d *DialogBox) ShowMenu(items []MenuItem) {
	d.menu = NewMenu(items)
	d.mode = DialogModeMenu
	d.Active = true
}

func (d *DialogBox) Close() {
	d.Active = false
	d.textScroll = nil
	d.menu = nil
}

func (d *DialogBox) Update() {
	if !d.Active {
		return
	}
	switch d.mode {
	case DialogModeText:
		d.textScroll.Update()
		if d.textScroll.IsDone() {
			d.Close()
		}
	case DialogModeMenu:
		d.menu.Update()
	}
}

func (d *DialogBox) Draw(screen *ebiten.Image) {
	if !d.Active {
		return
	}

	// black background behind box
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(DialogBoxX, DialogBoxY)
	screen.DrawImage(d.bg, op)

	// 9-slice border
	DrawDialogBox(screen, d.border, DialogBoxX, DialogBoxY, DialogBoxWidth, DialogBoxHeight)

	switch d.mode {
	case DialogModeText:
		d.textScroll.Draw(screen, d.font, DialogBoxX, DialogBoxY)
	case DialogModeMenu:
		d.menu.Draw(screen, DialogBoxX+DialogPadding, DialogBoxY+DialogPadding)
	}
}
