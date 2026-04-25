package location

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
	"os"

	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const TileSize = 16
const screenW = 320
const screenH = 240

type menuState int

const (
	menuClosed menuState = iota
	menuWipingIn
	menuOpen
	menuWipingOut
)

var PinkText = []float32{0.98, 0.653, 0.724, 1}
var MintText = []float32{0.641, 0.949, 0.678, 1}
var IcyBlueText = []float32{0.522, 0.682, 0.949, 1}
var BloodyText = []float32{0.6, 0.063, 0.047, 1}

type TriggerHandler func(name string)

type Location struct {
	ctx         *ctx.GameContext
	player      *entity.Player
	npcs        []*entity.NPC
	dialogBox   *ui.DialogBox
	triggers    []world.Trigger
	tileMap     *tiled.Map
	tilesets    []*ebiten.Image
	onTrigger   TriggerHandler
	returnScene scene.Scene
	exits       []scene.ExitConfig
	cameraX     float64
	cameraY     float64
	systemMenu   *ui.MenuStack
	menuState    menuState
	menuWipe     *transition.WipeTransition
	menuOverlay  *ebiten.Image
	shuttingDown bool
}

func NewLocation(
	c *ctx.GameContext,
	p *entity.Player,
	npcs []*entity.NPC,
	tm *tiled.Map,
	tilesets []*ebiten.Image,
	onTrigger TriggerHandler,
) *Location {
	overlay := ebiten.NewImage(screenW, screenH)
	overlay.Fill(color.Black)

	l := &Location{
		ctx:         c,
		player:      p,
		npcs:        npcs,
		dialogBox:   ui.NewDialogBox(c.Assets.Font, c.Assets.DialogBorder),
		triggers:    world.LoadTriggers(tm),
		tileMap:     tm,
		tilesets:    tilesets,
		onTrigger:   onTrigger,
		menuOverlay: overlay,
	}

	root := ui.NewMenu([]ui.MenuItem{
		{Label: "Characters", OnSelect: func() {
			l.systemMenu.Push(buildCharacterListMenu(c, l.systemMenu))
		}},
		{Label: "Items", OnSelect: func() {}},
		{Label: "Save", OnSelect: func() {}},
		{Label: "Exit Game", OnSelect: func() {
			c.Assets.TownBGM.Pause()
			c.Assets.GameStart.Rewind()
			c.Assets.GameStart.Play()
			l.shuttingDown = true
		}},
	})
	root.NavSound = c.Assets.MenuNav
	root.SelectSound = c.Assets.MenuSelect
	l.systemMenu = ui.NewMenuStackWithSounds(root, c.Assets.MenuCancel)

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
	if s.shuttingDown {
		if !s.ctx.Assets.GameStart.IsPlaying() {
			os.Exit(0)
		}
		return nil
	}

	switch s.menuState {
	case menuWipingIn:
		s.menuWipe.Update()
		if s.menuWipe.IsDone() {
			s.menuState = menuOpen
		}
		return nil
	case menuWipingOut:
		s.menuWipe.Update()
		if s.menuWipe.IsDone() {
			s.menuState = menuClosed
		}
		return nil
	case menuOpen:
		if (inpututil.IsKeyJustPressed(ebiten.KeyX) && len(s.systemMenu.Stack()) == 1) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
			s.ctx.Assets.MenuDown.Rewind()
			s.ctx.Assets.MenuDown.Play()
			s.menuWipe = transition.NewWipe(transition.WipeDown)
			s.menuState = menuWipingOut
			return nil
		}
		s.systemMenu.Update()
		return nil
	}

	if s.dialogBox.Active {
		s.dialogBox.Update()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		s.ctx.Assets.MenuUp.Rewind()
		s.ctx.Assets.MenuUp.Play()
		s.menuWipe = transition.NewWipe(transition.WipeUp)
		s.menuState = menuWipingIn
		return nil
	}

	s.player.Update(s.tileMap, s.npcs)

	for _, npc := range s.npcs {
		npc.Update(s.tileMap, s.player)
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
				s.ctx.ScSwitcher.SetScene(s.returnScene)
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

	switch s.menuState {
	case menuWipingIn:
		s.menuWipe.Draw(screen)
	case menuWipingOut:
		s.menuWipe.Draw(screen)
	case menuOpen:
		screen.DrawImage(s.menuOverlay, nil)
		s.systemMenu.Draw(screen, s.ctx.Assets.Font, 32, 32)
	}
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
