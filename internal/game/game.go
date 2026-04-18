package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
)

type Game struct {
	current scene.Scene
}

func New() *Game {
	return &Game{}
}

func (g *Game) SetScene(s scene.Scene) {
	g.current = s
}

func (g *Game) Update() error {
	return g.current.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}
