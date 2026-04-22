package location

import (
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

	t := &Town{firstEnter: true}

	thief.OnInteract = func() {
		t.dialogBox.ShowText(
			"Hey You!\n\nThere are some strange characters at the inn...\n\nuh...\nMaybe go do something about it?",
			nil,
			PinkText,
			a.VoiceOne,
		)
	}

	t.Location = *NewLocation(
		ss, a, p,
		[]*entity.NPC{thief},
		a.TownMap,
		a.TownTileset,
		func(name string) {
			if name == "inn_entrance" {
				returnX := t.player.X
				returnY := t.player.Y + 16
				t.player.Direction = 0
				inn := NewInn(ss, a, t, 240, 176, []scene.ExitConfig{
					{TriggerName: "exit", ReturnX: returnX, ReturnY: returnY},
				})
				ss.SetScene(inn)
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
