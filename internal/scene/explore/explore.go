package explore

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/interior"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const startX = 96
const startY = 200
const tileSize = 16

var pink = []float32{0.98, 0.653, 0.724, 1}

type ExploreScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	npcs          []*entity.NPC
	dialogBox     *ui.DialogBox
	triggers      []world.Trigger
	firstEnter    bool
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
		firstEnter:    true,
		npcs:          []*entity.NPC{thief},
		dialogBox:     d,
		triggers:      world.LoadTriggers(a.TownMap),
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

	trigger := world.CheckTrigger(s.triggers, s.player.X+8, s.player.Y+12)
	if trigger != nil && trigger.Name == "inn_entrance" {
		s.player.Y = s.player.Y + 16
		s.player.Direction = 0
		interiorScene := interior.New(s.sceneSwitcher, s.assets, s, 240, 176)
		s.sceneSwitcher.SetScene(interiorScene)
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
			"Hey You!\n\nThere are some strange characters at the inn...\n\nuh...\nMaybe go do something about it?",
			nil,
			pink,
			s.assets.VoiceOne,
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

func (s *ExploreScene) TransitionType() game.TransitionType {
	if s.firstEnter {
		return game.TransitionSpiral
	}
	return game.TransitionBox
}

func (s *ExploreScene) OnEnter() {
	s.firstEnter = false
	s.assets.TownBGM.Play()
	s.assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}

func (s *ExploreScene) SetPlayerPos(x, y float64) {
	s.player.X = x
	s.player.Y = y
}
