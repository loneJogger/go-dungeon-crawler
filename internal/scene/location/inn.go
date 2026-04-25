package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/battle"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Inn struct {
	*Location
}

func NewInn(c *ctx.GameContext, returnScene scene.Scene, startX, startY float64, exits []scene.ExitConfig) *Inn {
	p := entity.NewPlayer(startX, startY, c.Assets.PCSprite)
	p.Direction = 2

	devil := entity.NewNPC(128, 108, 0)
	devil.Image = c.Assets.NPCDevil
	devil.Wanders = true

	i := &Inn{}

	devil.OnInteract = func() {
		i.dialogBox.ShowText(
			"Hehehehehehehe",
			func() {
				c.Assets.TownBGM.Pause()
				c.Assets.BattleStart.Rewind()
				c.Assets.BattleStart.Play()
				c.Assets.BattleBGM.Rewind()
				c.Assets.BattleBGM.Play()
				c.ScSwitcher.SetScene(battle.New())
			},
			BloodyText,
			c.Assets.VoiceTwo,
		)
	}

	i.Location = NewLocation(c, p, []*entity.NPC{devil}, c.Assets.CaveMap, []*ebiten.Image{c.Assets.CaveTileset}, nil)
	i.returnScene = returnScene
	i.exits = exits

	return i
}

func (s *Inn) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Inn) TransitionType() transition.TransitionType {
	return transition.TransitionBox
}

func (s *Inn) OnEnter() {
	s.ctx.Assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
}

func (s *Inn) OnExit() {
	s.ctx.Assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
