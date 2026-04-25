package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

const townStartX = 96
const townStartY = 200

type Town struct {
	*Location
	firstEnter bool
}

func NewTown(c *ctx.GameContext) *Town {
	p := entity.NewPlayer(townStartX, townStartY, c.Assets.PCSprite)

	thief := entity.NewNPC(128, 160, 0)
	thief.Image = c.Assets.NPCThief
	thief.Wanders = true

	blackBelt := entity.NewNPC(272, 176, 1)
	blackBelt.Image = c.Assets.NPCBlackBelt
	blackBelt.Wanders = false

	t := &Town{firstEnter: true}

	thief.OnInteract = func() {
		t.dialogBox.ShowText(
			"Hey You!\n\nThere are some strange characters at the inn...\n\nuh...\nMaybe go do something about it?",
			nil,
			PinkText,
			c.Assets.VoiceOne,
		)
	}

	blackBelt.OnInteract = func() {
		t.dialogBox.ShowText(
			"... ... ... ... ... ... ... ...\n\nIf I was a sea gull, I would fly as far as I could!\n\nI would fly to far away places and sing for many people!\n\n",
			nil,
			MintText,
			c.Assets.VoiceTwo,
		)
	}

	t.Location = NewLocation(
		c, p,
		[]*entity.NPC{thief, blackBelt},
		c.Assets.TownMap,
		[]*ebiten.Image{c.Assets.TownTileset},
		func(name string) {
			switch name {
			case "inn_entrance":
				returnX := t.player.X
				returnY := t.player.Y + 16
				t.player.Direction = 0
				inn := NewInn(c, t, 240, 176, []scene.ExitConfig{
					{TriggerName: "inn_exit", ReturnX: returnX, ReturnY: returnY},
				})
				c.SS.SetScene(inn)
			case "west_shop_entrance":
				t.player.Direction = 0
				shop := NewShop(c, t, 128, 176, []scene.ExitConfig{
					{TriggerName: "west_shop_exit", ReturnX: 144, ReturnY: 144},
					{TriggerName: "east_shop_exit", ReturnX: 224, ReturnY: 144},
				})
				c.SS.SetScene(shop)
			case "east_shop_entrance":
				t.player.Direction = 0
				shop := NewShop(c, t, 528, 176, []scene.ExitConfig{
					{TriggerName: "west_shop_exit", ReturnX: 144, ReturnY: 144},
					{TriggerName: "east_shop_exit", ReturnX: 224, ReturnY: 144},
				})
				c.SS.SetScene(shop)
			}
		},
	)

	return t
}

func (s *Town) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Town) TransitionType() transition.TransitionType {
	if s.firstEnter {
		return transition.TransitionSpiral
	}
	return transition.TransitionBox
}

func (s *Town) OnEnter() {
	s.firstEnter = false
	s.ctx.Assets.TownBGM.Play()
	s.ctx.Assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
