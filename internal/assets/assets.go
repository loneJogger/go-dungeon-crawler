package assets

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Assets struct {
	DialogBorder *ebiten.Image
	TownMap      *tiled.Map
	TownTileset  *ebiten.Image
	PCSprite     *ebiten.Image
}

func LoadAssets() (*Assets, error) {
	dialogBorder, err := LoadImage("assets/ui/dialog_border.png")
	if err != nil {
		return nil, err
	}
	townMap, err := tiled.LoadFile("assets/tiledMaps/ff_town.tmx")
	if err != nil {
		return nil, err
	}
	townTileset, err := LoadImage("assets/tilesets/ff_town.png")
	if err != nil {
		return nil, err
	}
	pcSprite, err := LoadImage("assets/sprites/red_mage.png")
	if err != nil {
		return nil, err
	}
	return &Assets{
		DialogBorder: dialogBorder,
		TownMap:      townMap,
		TownTileset:  townTileset,
		PCSprite:     pcSprite,
	}, nil
}

func LoadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
