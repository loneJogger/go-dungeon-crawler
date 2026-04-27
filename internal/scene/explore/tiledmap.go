package explore

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
)

type Trigger struct {
	Name   string
	Bounds image.Rectangle
}

type TriggerHandler func(name string)

// returns a list of Triggers in a Tiled Map
func LoadTriggers(m *tiled.Map) []Trigger {
	var triggers []Trigger
	for _, group := range m.ObjectGroups {
		if group.Name != "triggers" {
			continue
		}
		for _, obj := range group.Objects {
			triggers = append(triggers, Trigger{
				Name: obj.Name,
				Bounds: image.Rect(
					int(obj.X),
					int(obj.Y),
					int(obj.X+obj.Width),
					int(obj.Y+obj.Height),
				),
			})
		}
	}
	return triggers
}

// returns a trigger at a given coordinate
func CheckTrigger(triggers []Trigger, x, y float64) *Trigger {
	pt := image.Pt(int(x), int(y))
	for i := range triggers {
		if pt.In(triggers[i].Bounds) {
			return &triggers[i]
		}
	}
	return nil
}

// checks a Tiled map for collision at a given coordinate
func IsCollison(m *tiled.Map, x, y float64) bool {
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

// draws a Tiled map to the screen
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

			x := float64((i%m.Width)*config.TileSize) - cx
			y := float64((i/m.Width)*config.TileSize) - cy

			id := int(tile.ID)
			tilesPerRow := img.Bounds().Dx() / config.TileSize
			tx := (id % tilesPerRow) * config.TileSize
			ty := (id / tilesPerRow) * config.TileSize

			src := img.SubImage(image.Rect(tx, ty, tx+config.TileSize, ty+config.TileSize)).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(src, op)
		}
	}
}
