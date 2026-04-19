package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 8

func DrawDialogBox(
	screen *ebiten.Image,
	border *ebiten.Image,
	x, y, width, height int,
) {
	// crop a tile from the strip by index
	tile := func(i int) *ebiten.Image {
		rect := image.Rect(i*tileSize, 0, (i+1)*tileSize, tileSize)
		return border.SubImage(rect).(*ebiten.Image)
	}

	// draw a single tile at a screen position
	draw := func(t *ebiten.Image, px, py int) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(px), float64(py))
		screen.DrawImage(t, op)
	}

	// corners (indices 0,2,6,8)
	draw(tile(0), x, y)                                // top-left
	draw(tile(2), x+width-tileSize, y)                 // top-right
	draw(tile(6), x, y+height-tileSize)                // bottom-left
	draw(tile(8), x+width-tileSize, y+height-tileSize) // bottom-right

	// top edges loop from x+8 to x+width-8 (index 1)
	for px := x + tileSize; px < x+width-tileSize; px += tileSize {
		draw(tile(1), px, y)
	}

	// left/right edges loop from y+8 to y+height-8 (indices 3,5)
	for py := y + tileSize; py < y+height-tileSize; py += tileSize {
		draw(tile(3), x, py)
		draw(tile(5), x+width-tileSize, py)
	}

	// center tiles fill the interior (index 4)
	for px := x + tileSize; px < x+width-tileSize; px += tileSize {
		for py := y + tileSize; py < y+height-tileSize; py += tileSize {
			draw(tile(4), px, py)
		}
	}

	// bottom edges loop (index 7)
	for px := x + tileSize; px < x+width-tileSize; px += tileSize {
		draw(tile(7), px, y+height-tileSize)
	}

}
