package explore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ExploreScene struct{}

func New() *ExploreScene {
	return &ExploreScene{}
}

func (s *ExploreScene) Update() error {
	return nil
}

func (s *ExploreScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Explore Scene")
}
