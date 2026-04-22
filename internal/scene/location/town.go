package location

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

const townStartX = 96
const townStartY = 200

type Town struct {
	Location
	firstEnter bool
}

func NewTown(ss scene.SceneSwitcher, a *assets.Assets) *Town {
	p := entity.NewPlayer(townStartX, townStartY, a.PCSprite)

	thief := entity.NewNPC(128, 160, 0)
	thief.Image = a.NPCThief
	thief.Wanders = true

	blackBelt := entity.NewNPC(272, 176, 1)
	blackBelt.Image = a.NPCBlackBelt
	blackBelt.Wanders = false

	t := &Town{firstEnter: true}

	thief.OnInteract = func() {
		t.dialogBox.ShowText(
			"Hey You!\n\nThere are some strange characters at the inn...\n\nuh...\nMaybe go do something about it?",
			nil,
			PinkText,
			a.VoiceOne,
		)
	}

	blackBelt.OnInteract = func() {
		t.dialogBox.ShowText(
			"... ... ... ... ... ... ... ...\n\nIf I was a sea gull, I would fly as far as I could!\n\nI would fly to far away places and sing for many people!\n\n",
			nil,
			MintText,
			a.VoiceTwo,
		)
	}

	t.Location = *NewLocation(
		ss, a, p,
		[]*entity.NPC{thief, blackBelt},
		a.TownMap,
		[]*ebiten.Image{a.TownTileset},
		func(name string) {
			switch name {
			case "inn_entrance":
				returnX := t.player.X
				returnY := t.player.Y + 16
				t.player.Direction = 0
				inn := NewInn(ss, a, t, 240, 176, []scene.ExitConfig{
					{TriggerName: "inn_exit", ReturnX: returnX, ReturnY: returnY},
				})
				ss.SetScene(inn)
			case "west_shop_entrance":
				t.player.Direction = 0
				shop := NewShop(ss, a, t, 128, 176, []scene.ExitConfig{
					{TriggerName: "west_shop_exit", ReturnX: 144, ReturnY: 144},
					{TriggerName: "east_shop_exit", ReturnX: 224, ReturnY: 144},
				})
				ss.SetScene(shop)
			case "east_shop_entrance":
				t.player.Direction = 0
				shop := NewShop(ss, a, t, 528, 176, []scene.ExitConfig{
					{TriggerName: "west_shop_exit", ReturnX: 144, ReturnY: 144},
					{TriggerName: "east_shop_exit", ReturnX: 224, ReturnY: 144},
				})
				ss.SetScene(shop)
			}
		},
	)

	return t
}

func (s *Town) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *Town) TransitionType() game.TransitionType {
	if s.firstEnter {
		return game.TransitionSpiral
	}
	return game.TransitionBox
}

func (s *Town) OnEnter() {
	s.firstEnter = false
	s.assets.TownBGM.Play()
	s.assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}
