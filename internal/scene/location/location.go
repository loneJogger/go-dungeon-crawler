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
const screenW = 320
const screenH = 240

var PinkText = []float32{0.98, 0.653, 0.724, 1}
var MintText = []float32{0.641, 0.949, 0.678, 1}

type TriggerHandler func(name string)

type Location struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	npcs          []*entity.NPC
	dialogBox     *ui.DialogBox
	triggers      []world.Trigger
	tileMap       *tiled.Map
	tilesets      []*ebiten.Image
	onTrigger     TriggerHandler
	returnScene   scene.Scene
	exits         []scene.ExitConfig
	cameraX       float64
	cameraY       float64
}

func NewLocation(
	ss scene.SceneSwitcher,
	a *assets.Assets,
	p *entity.Player,
	npcs []*entity.NPC,
	tm *tiled.Map,
	tilesets []*ebiten.Image,
	onTrigger TriggerHandler,
) *Location {
	l := &Location{
		sceneSwitcher: ss,
		assets:        a,
		player:        p,
		npcs:          npcs,
		dialogBox:     ui.NewDialogBox(a.Font, a.DialogBorder),
		triggers:      world.LoadTriggers(tm),
		tileMap:       tm,
		tilesets:      tilesets,
		onTrigger:     onTrigger,
	}
	l.updateCamera()
	return l
}

func (s *Location) updateCamera() {
	mapW := float64(s.tileMap.Width * s.tileMap.TileWidth)
	mapH := float64(s.tileMap.Height * s.tileMap.TileHeight)

	cx := s.player.X - screenW/2 + TileSize/2
	cy := s.player.Y - screenH/2 + TileSize/2

	if cx < 0 {
		cx = 0
	} else if cx > mapW-screenW {
		cx = mapW - screenW
	}
	if cy < 0 {
		cy = 0
	} else if cy > mapH-screenH {
		cy = mapH - screenH
	}

	s.cameraX = cx
	s.cameraY = cy
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

	s.updateCamera()

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		s.checkInteraction()
	}

	trigger := world.CheckTrigger(s.triggers, s.player.X+8, s.player.Y+12)
	if trigger != nil {
		for _, exit := range s.exits {
			if trigger.Name == exit.TriggerName {
				if pp, ok := s.returnScene.(scene.PlayerPositioner); ok {
					pp.SetPlayerPos(exit.ReturnX, exit.ReturnY)
				}
				s.sceneSwitcher.SetScene(s.returnScene)
				return nil
			}
		}
		if s.onTrigger != nil {
			s.onTrigger(trigger.Name)
		}
	}

	return nil
}

func (s *Location) Draw(screen *ebiten.Image) {
	world.DrawMap(screen, s.tileMap, s.tilesets, s.cameraX, s.cameraY)
	for _, npc := range s.npcs {
		npc.OffsetX, npc.OffsetY = s.cameraX, s.cameraY
		npc.Draw(screen)
	}
	s.player.OffsetX, s.player.OffsetY = s.cameraX, s.cameraY
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
