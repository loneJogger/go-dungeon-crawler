package world

import "github.com/lafriks/go-tiled"

func IsSolid(m *tiled.Map, x, y float64) bool {
	tileX := int(x) / 16
	tileY := int(y) / 16
	for _, layer := range m.Layers {
		if layer.Name == "collision" {
			tile := layer.Tiles[tileY*m.Width+tileX]
			return !tile.IsNil()
		}
	}
	return false
}
