package assets

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/lafriks/go-tiled"
)

type Assets struct {
	// sprites
	PCSprite     *ebiten.Image
	DialogBorder *ebiten.Image

	// tilesets
	TownMap     *tiled.Map
	TownTileset *ebiten.Image

	AudioContext *audio.Context
	// sfx
	MenuNav    *audio.Player
	MenuSelect *audio.Player
}

func LoadAssets() (*Assets, error) {
	audioContext := audio.NewContext(44100)

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
	assets := &Assets{
		DialogBorder: dialogBorder,
		TownMap:      townMap,
		TownTileset:  townTileset,
		PCSprite:     pcSprite,
		AudioContext: audioContext,
	}

	menuNav, err := assets.LoadAudio("assets/sfx/menuNav.wav")
	if err != nil {
		return nil, err
	}
	menuSelect, err := assets.LoadAudio("assets/sfx/menuSelect.wav")
	if err != nil {
		return nil, err
	}

	assets.MenuNav = menuNav
	assets.MenuSelect = menuSelect

	return assets, nil
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

func (a *Assets) LoadAudio(path string) (*audio.Player, error) {
	audioFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer audioFile.Close()
	audioDecoded, err := wav.DecodeWithoutResampling(audioFile)
	if err != nil {
		return nil, err
	}
	sound, err := a.AudioContext.NewPlayer(audioDecoded)
	if err != nil {
		return nil, err
	}
	return sound, nil
}
