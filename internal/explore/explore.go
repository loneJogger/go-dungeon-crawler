package explore

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/rooms"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

type menuState int

const (
	menuClosed menuState = iota
	menuWipingIn
	menuOpen
	menuWipingOut
)

type Explore struct {
	ctx          *ctx.GameContext
	shuttingDown bool

	systemMenu  *ui.MenuStack
	menuState   menuState
	menuWipe    *transition.WipeTransition
	menuOverlay *ebiten.Image

	currentRoom    *room.Room
	pendingRoom    *room.Room
	pendingRoomX   float64
	pendingRoomY   float64
	roomTransition transition.Transition
}

func New(c *ctx.GameContext) *Explore {
	overlay := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	overlay.Fill(color.Black)

	e := &Explore{
		ctx:         c,
		menuOverlay: overlay,
	}

	rootMenu := ui.NewMenu([]ui.MenuItem{
		{Label: "Characters", OnSelect: func() {
			e.systemMenu.Push(BuildCharacterListMenu(c, e.systemMenu))
		}},
		{Label: "Items", OnSelect: func() {}},
		{Label: "Save", OnSelect: func() {}},
		{Label: "Exit Game", OnSelect: func() {
			c.Assets.TownBGM.Pause()
			c.Assets.GameStart.Rewind()
			c.Assets.GameStart.Play()
			e.shuttingDown = true
		}},
	})
	e.systemMenu = ui.NewMenuStack(rootMenu, ui.MenuSounds{
		Nav:    c.Assets.MenuNav,
		Select: c.Assets.MenuSelect,
		Cancel: c.Assets.MenuCancel,
	})

	town := rooms.NewTownRoom(c, e.switchRoom)
	town.SetPlayerPos(rooms.TownStartX, rooms.TownStartY)
	town.UpdateCamera()
	e.currentRoom = town

	return e
}

func (e *Explore) switchRoom(r *room.Room, tt transition.TransitionType, startX, startY float64) {
	if e.currentRoom != nil && e.currentRoom.OnExit != nil {
		e.currentRoom.OnExit()
	}
	e.pendingRoom = r
	e.pendingRoomX = startX
	e.pendingRoomY = startY
	switch tt {
	case transition.TransitionBox:
		e.roomTransition = transition.NewBox(transition.Closing)
	default:
		e.roomTransition = transition.New(transition.Closing)
	}
}

func (e *Explore) TransitionPhase() transition.Phase { return transition.Opening }
func (e *Explore) TransitionType() transition.TransitionType {
	return transition.TransitionSpiral
}

func (e *Explore) OnEnter() {
	e.ctx.Assets.TownBGM.Play()
	e.ctx.Assets.TownBGM.SetVolume(assets.BgmWorldVolume)
}

func (e *Explore) SetPlayerPos(x, y float64) {
	e.currentRoom.SetPlayerPos(x, y)
}

func (e *Explore) Update() error {
	if e.shuttingDown {
		if !e.ctx.Assets.GameStart.IsPlaying() {
			os.Exit(0)
		}
		return nil
	}

	if e.roomTransition != nil {
		e.roomTransition.Update()
		if e.roomTransition.IsFullyBlack() && e.pendingRoom != nil {
			e.currentRoom = e.pendingRoom
			e.currentRoom.SetPlayerPos(e.pendingRoomX, e.pendingRoomY)
			e.currentRoom.UpdateCamera()
			if e.currentRoom.OnEnter != nil {
				e.currentRoom.OnEnter()
			}
			e.pendingRoom = nil
		}
		if e.roomTransition.IsDone() {
			e.roomTransition = nil
		}
		return nil
	}

	switch e.menuState {
	case menuWipingIn:
		e.menuWipe.Update()
		if e.menuWipe.IsDone() {
			e.menuState = menuOpen
		}
		return nil
	case menuWipingOut:
		e.menuWipe.Update()
		if e.menuWipe.IsDone() {
			e.menuState = menuClosed
		}
		return nil
	case menuOpen:
		if (inpututil.IsKeyJustPressed(ebiten.KeyX) && len(e.systemMenu.Stack()) == 1) ||
			inpututil.IsKeyJustPressed(ebiten.KeyA) {
			e.ctx.Assets.MenuDown.Rewind()
			e.ctx.Assets.MenuDown.Play()
			e.menuWipe = transition.NewWipe(transition.WipeDown)
			e.menuState = menuWipingOut
			return nil
		}
		e.systemMenu.Update()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		e.ctx.Assets.MenuUp.Rewind()
		e.ctx.Assets.MenuUp.Play()
		e.menuWipe = transition.NewWipe(transition.WipeUp)
		e.menuState = menuWipingIn
		return nil
	}

	e.currentRoom.Update()
	return nil
}

func (e *Explore) Draw(screen *ebiten.Image) {
	e.currentRoom.Draw(screen)

	switch e.menuState {
	case menuWipingIn:
		e.menuWipe.Draw(screen)
	case menuWipingOut:
		e.menuWipe.Draw(screen)
	case menuOpen:
		screen.DrawImage(e.menuOverlay, nil)
		e.systemMenu.Draw(screen, e.ctx.Assets.Font, 32, 32)
	}

	if e.roomTransition != nil {
		e.roomTransition.Draw(screen)
	}
}
