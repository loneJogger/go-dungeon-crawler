package location

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Inn struct {
	Location
	returnScene scene.Scene
	exits       []scene.ExitConfig
}

func NewInn(ss scene.SceneSwitcher, a *assets.Assets, returnScene scene.Scene, startX, startY float64, exits []scene.ExitConfig) *Inn {
	p := entity.NewPlayer(startX, startY, a.PCSprite)
	p.Direction = 2

	i := &Inn{
		returnScene: returnScene,
		exits:       exits,
	}

	i.Location = *NewLocation(
		ss, a, p,
		nil,
		a.CaveMap,
		a.CaveTileset,
		func(name string) {
			for _, exit := range i.exits {
				if exit.TriggerName == name {
					if pp, ok := returnScene.(scene.PlayerPositioner); ok {
						pp.SetPlayerPos(exit.ReturnX, exit.ReturnY)
					}
					ss.SetScene(returnScene)
					return
				}
			}
		},
	)

	return i
}

func (s *Inn) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Inn) TransitionType() game.TransitionType {
	return game.TransitionBox
}

func (s *Inn) OnEnter() {
	s.assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
}

func (s *Inn) OnExit() {
	s.assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
