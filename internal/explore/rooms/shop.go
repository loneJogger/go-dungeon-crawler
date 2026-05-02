package rooms

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

func NewShopRoom(c *ctx.GameContext, switchRoom SwitchFn, westExitX, westExitY, eastExitX, eastExitY float64) *room.Room {
	p := entity.NewPlayer(0, 0, c.Assets.PCSprite)
	p.Direction = 2

	weaponMerchant := entity.NewNPC(144, 112, 0)
	weaponMerchant.Image = c.Assets.NPCThief
	weaponMerchant.Wanders = false

	itemMerchant := entity.NewNPC(464, 112, 0)
	itemMerchant.Image = c.Assets.NPCBlackBelt
	itemMerchant.Wanders = false

	r := room.New(
		c, p,
		[]*entity.NPC{weaponMerchant, itemMerchant},
		c.Assets.ShopMap,
		[]*ebiten.Image{c.Assets.CaveTileset, c.Assets.TownTileset},
	)

	weaponMerchant.OnInteract = func() {
		r.ShowDialog(
			"Come see me once someone invents money.",
			nil,
			config.IcyBlueText,
			c.Assets.VoiceOne,
		)
	}

	itemMerchant.OnInteract = func() {
		r.ShowDialog(
			"Come see me once someone invents money.",
			nil,
			config.IcyBlueText,
			c.Assets.VoiceOne,
		)
	}

	r.OnTrigger = func(name string) {
		switch name {
		case "west_shop_exit":
			switchRoom(NewTownRoom(c, switchRoom), transition.TransitionBox, westExitX, westExitY)
		case "east_shop_exit":
			switchRoom(NewTownRoom(c, switchRoom), transition.TransitionBox, eastExitX, eastExitY)
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
