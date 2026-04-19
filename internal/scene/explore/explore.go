package explore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

type ExploreScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
}

func New(ss scene.SceneSwitcher, a *assets.Assets) *ExploreScene {
	s := &ExploreScene{sceneSwitcher: ss, assets: a}
	return s
}

func (s *ExploreScene) Update() error {
	return nil
}

func (s *ExploreScene) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.assets.TownMap, s.assets.TownTileset)
}
