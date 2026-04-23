package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Shop struct {
	Location
}

func NewShop(ss scene.SceneSwitcher, a *assets.Assets, returnScene scene.Scene, startX, startY float64, exits []scene.ExitConfig) *Shop {
	p := entity.NewPlayer(startX, startY, a.PCSprite)
	p.Direction = 2

	weaponMerchant := entity.NewNPC(144, 112, 0)
	weaponMerchant.Image = a.NPCThief
	weaponMerchant.Wanders = false

	itemMerchant := entity.NewNPC(464, 112, 0)
	itemMerchant.Image = a.NPCBlackBelt
	itemMerchant.Wanders = false

	sh := &Shop{}

	weaponMerchant.OnInteract = func() {
		sh.dialogBox.ShowText(
			"Come see me once someone invents money.",
			nil,
			IcyBlueText,
			a.VoiceOne,
		)
	}

	itemMerchant.OnInteract = func() {
		sh.dialogBox.ShowText(
			"Come see me once someone invents money.",
			nil,
			IcyBlueText,
			a.VoiceOne,
		)
	}

	sh.Location = *NewLocation(
		ss,
		a,
		p,
		[]*entity.NPC{weaponMerchant, itemMerchant},
		a.ShopMap,
		[]*ebiten.Image{a.CaveTileset, a.TownTileset},
		nil,
	)
	sh.returnScene = returnScene
	sh.exits = exits

	return sh
}

func (s *Shop) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Shop) TransitionType() game.TransitionType {
	return game.TransitionBox
}

func (s *Shop) OnEnter() {
	s.assets.TownBGM.SetVolume(assets.BgmInteriorVolume)
}

func (s *Shop) OnExit() {
	s.assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
