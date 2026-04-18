package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TitleScene struct{}

func New() *TitleScene { return &TitleScene{} }

func (s *TitleScene) Update() error { return nil }

func (s *TitleScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Explore Scene")
}
