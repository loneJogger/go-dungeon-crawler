package world

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

const tileSize = 16

func DrawMap(screen *ebiten.Image, m *tiled.Map, tilesets []*ebiten.Image, cx, cy float64) {
	for _, layer := range m.Layers {
		for i, tile := range layer.Tiles {
			if tile.IsNil() {
				continue
			}

			tsIdx := 0
			for j, ts := range m.Tilesets {
				if ts == tile.Tileset {
					tsIdx = j
					break
				}
			}
			if tsIdx >= len(tilesets) {
				continue
			}
			img := tilesets[tsIdx]

			x := float64((i%m.Width)*tileSize) - cx
			y := float64((i/m.Width)*tileSize) - cy

			id := int(tile.ID)
			tilesPerRow := img.Bounds().Dx() / tileSize
			tx := (id % tilesPerRow) * tileSize
			ty := (id / tilesPerRow) * tileSize

			src := img.SubImage(image.Rect(tx, ty, tx+tileSize, ty+tileSize)).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(src, op)
		}
	}
}
