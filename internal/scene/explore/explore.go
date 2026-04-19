package explore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const startX = 96
const startY = 200

type ExploreScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
}

func New(ss scene.SceneSwitcher, a *assets.Assets) *ExploreScene {
	p := entity.NewPlayer(startX, startY)
	s := &ExploreScene{sceneSwitcher: ss, assets: a, player: p}
	return s
}

func (s *ExploreScene) Update() error {
	s.player.Update(s.assets.TownMap)
	return nil
}

func (s *ExploreScene) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.assets.TownMap, s.assets.TownTileset)
	s.player.Draw(screen, s.assets.PCSprite)
}

func (s *ExploreScene) TransitionPhase() transition.Phase {
	return transition.Opening
}
