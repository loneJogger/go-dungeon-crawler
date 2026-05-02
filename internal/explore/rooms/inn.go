package rooms

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/battle"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

func NewInnRoom(c *ctx.GameContext, switchRoom SwitchFn, returnX, returnY float64) *room.Room {
	p := entity.NewPlayer(0, 0, c.Assets.PCSprite)
	p.Direction = 2

	devil := entity.NewNPC(128, 108, 0)
	devil.Image = c.Assets.NPCDevil
	devil.Wanders = true

	r := room.New(c, p, []*entity.NPC{devil}, c.Assets.CaveMap, []*ebiten.Image{c.Assets.CaveTileset})

	devil.OnInteract = func() {
		r.ShowDialog(
			"{COLOR BLOODY}Hehehehehehehe",
			func() {
				c.Assets.TownBGM.Pause()
				c.Assets.BattleStart.Rewind()
				c.Assets.BattleStart.Play()
				c.Assets.BattleBGM.Rewind()
				c.Assets.BattleBGM.Play()
				c.ScSwitcher.SetScene(battle.New())
			},
			c.Assets.VoiceTwo,
		)
	}

	r.OnTrigger = func(name string) {
		if name == "inn_exit" {
			switchRoom(NewTownRoom(c, switchRoom), transition.TransitionBox, returnX, returnY)
		}
	}

	r.OnEnter = func() {
		c.Assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
	}

	r.OnExit = func() {
		c.Assets.TownBGM.SetVolume(assets.BgmWorldVolume)
	}

	return r
}
