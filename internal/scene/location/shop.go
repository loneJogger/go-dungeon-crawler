package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Shop struct {
	*Location
}

func NewShop(c *ctx.GameContext, returnScene scene.Scene, startX, startY float64, exits []scene.ExitConfig) *Shop {
	p := entity.NewPlayer(startX, startY, c.Assets.PCSprite)
	p.Direction = 2

	weaponMerchant := entity.NewNPC(144, 112, 0)
	weaponMerchant.Image = c.Assets.NPCThief
	weaponMerchant.Wanders = false

	itemMerchant := entity.NewNPC(464, 112, 0)
	itemMerchant.Image = c.Assets.NPCBlackBelt
	itemMerchant.Wanders = false

	sh := &Shop{}

	weaponMerchant.OnInteract = func() {
		sh.dialogBox.ShowText(
			"Come see me once someone invents money.",
			nil,
			IcyBlueText,
			c.Assets.VoiceOne,
		)
	}

	itemMerchant.OnInteract = func() {
		sh.dialogBox.ShowText(
			"Come see me once someone invents money.",
			nil,
			IcyBlueText,
			c.Assets.VoiceOne,
		)
	}

	sh.Location = NewLocation(
		c,
		p,
		[]*entity.NPC{weaponMerchant, itemMerchant},
		c.Assets.ShopMap,
		[]*ebiten.Image{c.Assets.CaveTileset, c.Assets.TownTileset},
		nil,
	)
	sh.returnScene = returnScene
	sh.exits = exits

	return sh
}

func (s *Shop) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Shop) TransitionType() transition.TransitionType {
	return transition.TransitionBox
}

func (s *Shop) OnEnter() {
	s.ctx.Assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
}

func (s *Shop) OnExit() {
	s.ctx.Assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
