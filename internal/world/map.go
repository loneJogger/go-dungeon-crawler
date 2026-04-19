package world

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

const tileSize = 16

func DrawMap(screen *ebiten.Image, m *tiled.Map, tileset *ebiten.Image) {
	for _, layer := range m.Layers {
		for i, tile := range layer.Tiles {
			if tile.IsNil() {
				continue
			}

			// position on screen
			x := (i % m.Width) * tileSize
			y := (i / m.Width) * tileSize

			// position in tileset
			id := int(tile.ID)
			tilesPerRow := tileset.Bounds().Dx() / tileSize
			tx := (id % tilesPerRow) * tileSize
			ty := (id / tilesPerRow) * tileSize

			src := tileset.SubImage(image.Rect(tx, ty, tx+tileSize, ty+tileSize)).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(src, op)
		}
	}
}
