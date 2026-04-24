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
const BgmWorldVolume = 0.5
const BgmInteriorVolume = 0.25

type Assets struct {
	// sprites
	PCSprite     *ebiten.Image
	NPCThief     *ebiten.Image
	NPCBlackBelt *ebiten.Image
	NPCDevil     *ebiten.Image
	Font         *ebiten.Image
	DialogBorder *ebiten.Image

	// tilesets
	TownMap     *tiled.Map
	TownTileset *ebiten.Image
	CaveMap     *tiled.Map
	CaveTileset *ebiten.Image
	ShopMap     *tiled.Map

	// sound
	AudioContext *audio.Context
	// sfx
	MenuNav     *audio.Player
	MenuSelect  *audio.Player
	MenuCancel  *audio.Player
	GameStart   *audio.Player
	BattleStart *audio.Player
	VoiceOne    *audio.Player
	VoiceTwo    *audio.Player
	// bgm
	TitleBGM  *audio.Player
	TownBGM   *audio.Player
	BattleBGM *audio.Player
}

type imageEntry struct {
	dest *(*ebiten.Image)
	path string
}

type soundEntry struct {
	dest   *(*audio.Player)
	path   string
	volume float64
}

type bgmEntry struct {
	dest   *(*audio.Player)
	path   string
	volume float64
}

func LoadAssets() (*Assets, error) {
	audioContext := audio.NewContext(44100)
	a := &Assets{AudioContext: audioContext}

	images := []imageEntry{
		{&a.DialogBorder, "assets/ui/dialog_border.png"},
		{&a.Font, "assets/ui/8BitFont.png"},
		{&a.TownTileset, "assets/tilesets/ff_town.png"},
		{&a.CaveTileset, "assets/tilesets/ff_cave.png"},
		{&a.PCSprite, "assets/sprites/red_mage.png"},
		{&a.NPCThief, "assets/sprites/thief.png"},
		{&a.NPCBlackBelt, "assets/sprites/black_belt.png"},
		{&a.NPCDevil, "assets/sprites/devil.png"},
	}
	for _, e := range images {
		img, err := LoadImage(e.path)
		if err != nil {
			return nil, err
		}
		*e.dest = img
	}

	maps := []struct {
		dest *(*tiled.Map)
		path string
	}{
		{&a.TownMap, "assets/tiledMaps/ff_town.tmx"},
		{&a.CaveMap, "assets/tiledMaps/ff_cave.tmx"},
		{&a.ShopMap, "assets/tiledMaps/shop.tmx"},
	}
	for _, e := range maps {
		m, err := tiled.LoadFile(e.path)
		if err != nil {
			return nil, err
		}
		*e.dest = m
	}

	sounds := []soundEntry{
		{&a.MenuNav, "assets/sfx/menuNav.wav", sfxSoftVolume},
		{&a.MenuSelect, "assets/sfx/menuSelect.wav", sfxSoftVolume},
		{&a.MenuCancel, "assets/sfx/menuCancel.wav", sfxSoftVolume},
		{&a.GameStart, "assets/sfx/gameStart.wav", sfxLoudVolume},
		{&a.BattleStart, "assets/sfx/battleStart.wav", sfxLoudVolume},
		{&a.VoiceOne, "assets/sfx/voice1.wav", BgmWorldVolume},
		{&a.VoiceTwo, "assets/sfx/voice2.wav", BgmWorldVolume},
	}
	for _, e := range sounds {
		p, err := a.LoadSound(e.path)
		if err != nil {
			return nil, err
		}
		p.SetVolume(e.volume)
		*e.dest = p
	}

	bgms := []bgmEntry{
		{&a.TitleBGM, "assets/bgMusic/title_theme.ogg", bgmTitleVolume},
		{&a.TownBGM, "assets/bgMusic/town_theme.ogg", BgmWorldVolume},
		{&a.BattleBGM, "assets/bgMusic/battle_theme.ogg", bgmTitleVolume},
	}
	for _, e := range bgms {
		p, err := a.LoadBGM(e.path)
		if err != nil {
			return nil, err
		}
		p.SetVolume(e.volume)
		*e.dest = p
	}

	return a, nil
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
