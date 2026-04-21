package assets

import (
	"bytes"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/lafriks/go-tiled"
)

const sfxLoudVolume = 0.9
const sfxSoftVolume = 0.8
const bgmTitleVolume = 0.7
const bgmWorldVolume = 0.5
const bgmInteriorVolume = 0.25

type Assets struct {
	// sprites
	PCSprite     *ebiten.Image
	DialogBorder *ebiten.Image

	// tilesets
	TownMap     *tiled.Map
	TownTileset *ebiten.Image

	// sound
	AudioContext *audio.Context
	// sfx
	MenuNav    *audio.Player
	MenuSelect *audio.Player
	GameStart  *audio.Player
	// bgm
	TitleBGM *audio.Player
	TownBGM  *audio.Player
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

	menuNav, err := assets.LoadSound("assets/sfx/menuNav.wav")
	if err != nil {
		return nil, err
	}
	menuNav.SetVolume(sfxSoftVolume)
	menuSelect, err := assets.LoadSound("assets/sfx/menuSelect.wav")
	if err != nil {
		return nil, err
	}
	menuSelect.SetVolume(sfxSoftVolume)
	gameStart, err := assets.LoadSound("assets/sfx/gameStart.wav")
	if err != nil {
		return nil, err
	}
	gameStart.SetVolume(sfxLoudVolume)
	titleBgm, err := assets.LoadBGM("assets/bgMusic/title_theme.ogg")
	if err != nil {
		return nil, err
	}
	titleBgm.SetVolume(bgmTitleVolume)
	townBgm, err := assets.LoadBGM("assets/bgMusic/town_theme.ogg")
	if err != nil {
		return nil, err
	}
	townBgm.SetVolume(bgmWorldVolume)

	assets.MenuNav = menuNav
	assets.MenuSelect = menuSelect
	assets.GameStart = gameStart
	assets.TitleBGM = titleBgm
	assets.TownBGM = townBgm

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

func (a *Assets) LoadSound(path string) (*audio.Player, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	audioDecoded, err := wav.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	player, err := a.AudioContext.NewPlayer(audioDecoded)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (a *Assets) LoadBGM(path string) (*audio.Player, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	bgmDecoded, err := vorbis.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	loop := audio.NewInfiniteLoop(bgmDecoded, bgmDecoded.Length())
	player, err := a.AudioContext.NewPlayer(loop)
	if err != nil {
		return nil, err
	}
	return player, nil
}
