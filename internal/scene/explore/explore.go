package explore

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const startX = 96
const startY = 200
const tileSize = 16

type ExploreScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	npcs          []*entity.NPC
	dialogBox     *ui.DialogBox
}

func New(ss scene.SceneSwitcher, a *assets.Assets) *ExploreScene {
	p := entity.NewPlayer(startX, startY, a.PCSprite)
	d := ui.NewDialogBox(a.Font, a.DialogBorder)

	thief := entity.NewNPC(128, 160, 0, entity.InteractionDialog)
	thief.Image = a.NPCThief

	s := &ExploreScene{
		sceneSwitcher: ss,
		assets:        a,
		player:        p,
		npcs:          []*entity.NPC{thief},
		dialogBox:     d,
	}
	return s
}

func (s *ExploreScene) Update() error {
	if s.dialogBox.Active {
		s.dialogBox.Update()
		return nil
	}

	s.player.Update(s.assets.TownMap, s.npcs)

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		s.checkInteraction()
	}

	return nil
}

func (s *ExploreScene) checkInteraction() {
	// calculate the tile directly in front of the player
	checkX := s.player.X + tileSize/2
	checkY := s.player.Y + tileSize/2

	switch s.player.Direction {
	case 0: // down
		checkY += tileSize
	case 2: // up
		checkY -= tileSize
	case 1:
		if s.player.FacingRight {
			checkX += tileSize
		} else {
			checkX -= tileSize
		}
	}

	pt := image.Pt(int(checkX), int(checkY))

	for _, npc := range s.npcs {
		if pt.In(npc.Bounds()) {
			s.triggerInteraction(npc)
			return
		}
	}
}

func (s *ExploreScene) triggerInteraction(npc *entity.NPC) {
	switch npc.Interaction {
	case entity.InteractionDialog:
		s.dialogBox.ShowText(
			"Hello, traveler!\n\nI wouldn't go into those woods alone.\n\n",
			nil,
			[]float32{1, 0.8, 0, 1},
		)
	case entity.InteractionBattle:
		// TODO: trigger battle scene
	}
}

func (s *ExploreScene) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.assets.TownMap, s.assets.TownTileset)
	for _, npc := range s.npcs {
		npc.Draw(screen)
	}
	s.player.Draw(screen)
	s.dialogBox.Draw(screen)
}

func (s *ExploreScene) TransitionPhase() transition.Phase {
	return transition.Opening
}

func (s *ExploreScene) OnEnter() {
	s.assets.TownBGM.Play()
}
