package interior

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

type InteriorScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	triggers      []world.Trigger
	dialogBox     *ui.DialogBox
	returnScene   scene.Scene
}

func New(ss scene.SceneSwitcher, a *assets.Assets, returnScene scene.Scene, startX, startY float64) *InteriorScene {
	s := &InteriorScene{
		sceneSwitcher: ss,
		assets:        a,
		player:        entity.NewPlayer(startX, startY, a.PCSprite),
		triggers:      world.LoadTriggers(a.CaveMap),
		dialogBox:     ui.NewDialogBox(a.Font, a.DialogBorder),
		returnScene:   returnScene,
	}
	s.player.Direction = 2
	return s
}

func (s *InteriorScene) Update() error {
	if s.dialogBox.Active {
		s.dialogBox.Update()
		return nil
	}

	s.player.Update(s.assets.CaveMap, nil)

	trigger := world.CheckTrigger(s.triggers, s.player.X+8, s.player.Y+12)
	if trigger != nil && trigger.Name == "exit" {
		s.sceneSwitcher.SetScene(s.returnScene)
	}

	return nil
}

func (s *InteriorScene) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.assets.CaveMap, s.assets.CaveTileset)
	s.player.Draw(screen)
	s.dialogBox.Draw(screen)
}

func (s *InteriorScene) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *InteriorScene) TransitionType() game.TransitionType {
	return game.TransitionBox
}

func (s *InteriorScene) OnEnter() {
	s.assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
}

func (s *InteriorScene) OnExit() {
	s.assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
