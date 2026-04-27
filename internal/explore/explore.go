// a package for all components having to do with the explore side of the game
package explore

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
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
	systemMenu   *ui.MenuStack
	menuState    menuState
	menuWipe     *transition.WipeTransition
	menuOverlay  *ebiten.Image
	currentRoom  *room.Room
	shuttingDown bool
}

func NewExplore(
	ctx *ctx.GameContext,
	r *room.Room,
) *Explore {
	overlay := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	overlay.Fill(color.Black)

	e := &Explore{
		ctx: ctx,
	}

	rootMenu := ui.NewMenu([]ui.MenuItem{
		{Label: "Characters", OnSelect: func() {
			e.systemMenu.Push(BuildCharacterListMenu(ctx, e.systemMenu))
		}},
		{Label: "Items", OnSelect: func() {}},
		{Label: "Save", OnSelect: func() {}},
		{Label: "Exit Game", OnSelect: func() {
			ctx.Assets.TownBGM.Pause()
			ctx.Assets.GameStart.Rewind()
			ctx.Assets.GameStart.Play()
			e.shuttingDown = true
		}},
	})
	e.systemMenu = ui.NewMenuStack(rootMenu, ui.MenuSounds{
		Nav:    ctx.Assets.MenuNav,
		Select: ctx.Assets.MenuSelect,
		Cancel: ctx.Assets.MenuCancel,
	})

	return nil
}
