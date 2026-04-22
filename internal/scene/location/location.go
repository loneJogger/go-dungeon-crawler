package location

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const TileSize = 16

var PinkText = []float32{0.98, 0.653, 0.724, 1}

type TriggerHandler func(name string)

type Location struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	npcs          []*entity.NPC
	dialogBox     *ui.DialogBox
	triggers      []world.Trigger
	tileMap       *tiled.Map
	tileset       *ebiten.Image
	onTrigger     TriggerHandler
}

func NewLocation(
	ss scene.SceneSwitcher,
	a *assets.Assets,
	p *entity.Player,
	npcs []*entity.NPC,
	tm *tiled.Map,
	ts *ebiten.Image,
	onTrigger TriggerHandler,
) *Location {
	return &Location{
		sceneSwitcher: ss,
		assets:        a,
		player:        p,
		npcs:          npcs,
		dialogBox:     ui.NewDialogBox(a.Font, a.DialogBorder),
		triggers:      world.LoadTriggers(tm),
		tileMap:       tm,
		tileset:       ts,
		onTrigger:     onTrigger,
	}
}

func (s *Location) Update() error {
	if s.dialogBox.Active {
		s.dialogBox.Update()
		return nil
	}

	s.player.Update(s.tileMap, s.npcs)

	for _, npc := range s.npcs {
		npc.Update(s.tileMap)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		s.checkInteraction()
	}

	trigger := world.CheckTrigger(s.triggers, s.player.X+8, s.player.Y+12)
	if trigger != nil && s.onTrigger != nil {
		s.onTrigger(trigger.Name)
	}

	return nil
}

func (s *Location) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.tileMap, s.tileset)
	for _, npc := range s.npcs {
		npc.Draw(screen)
	}
	s.player.Draw(screen)
	s.dialogBox.Draw(screen)
}

func (s *Location) SetPlayerPos(x, y float64) {
	s.player.X = x
	s.player.Y = y
}

func (s *Location) checkInteraction() {
	checkX := s.player.X + TileSize/2
	checkY := s.player.Y + TileSize/2

	switch s.player.Direction {
	case 0:
		checkY += TileSize
	case 2:
		checkY -= TileSize
	case 1:
		if s.player.FacingRight {
			checkX += TileSize
		} else {
			checkX -= TileSize
		}
	}

	pt := image.Pt(int(checkX), int(checkY))

	for _, npc := range s.npcs {
		if pt.In(npc.Bounds()) && npc.OnInteract != nil {
			npc.OnInteract()
			return
		}
	}
}
