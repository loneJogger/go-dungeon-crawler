package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/battle"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Inn struct {
	Location
}

func NewInn(ss scene.SceneSwitcher, a *assets.Assets, returnScene scene.Scene, startX, startY float64, exits []scene.ExitConfig) *Inn {
	p := entity.NewPlayer(startX, startY, a.PCSprite)
	p.Direction = 2

	devil := entity.NewNPC(128, 64, 0)
	devil.Image = a.NPCDevil
	// devil.Wanders = true

	i := &Inn{}

	devil.OnInteract = func() {
		i.dialogBox.ShowText(
			"Hehehehehehehe",
			func() {
				a.TownBGM.Pause()
				a.BattleStart.Rewind()
				a.BattleStart.Play()
				a.BattleBGM.Rewind()
				a.BattleBGM.Play()
				ss.SetScene(battle.New())

			},
			BloodyText,
			a.VoiceTwo,
		)
	}

	i.Location = *NewLocation(ss, a, p, []*entity.NPC{devil}, a.CaveMap, []*ebiten.Image{a.CaveTileset}, nil)
	i.returnScene = returnScene
	i.exits = exits

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
