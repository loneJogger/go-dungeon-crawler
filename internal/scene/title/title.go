package title

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/location"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

const menuX = 64
const menuY = 144

type TitleScene struct {
	ctx  *ctx.GameContext
	menu ui.Menu
}

func New(c *ctx.GameContext) *TitleScene {
	s := &TitleScene{ctx: c}
	s.menu = *ui.NewMenu([]ui.MenuItem{
		{Label: "New Game", OnSelect: func() {
			c.Assets.GameStart.Rewind()
			c.Assets.GameStart.Play()
			c.ScSwitcher.SetScene(location.NewTown(c))
		}},
		{Label: "Continue", OnSelect: func() { /* TODO */ }},
		{Label: "Exit", OnSelect: func() { os.Exit(0) }},
	})
	s.menu.NavSound = c.Assets.MenuNav
	s.menu.SelectSound = c.Assets.MenuSelect
	c.Assets.TitleBGM.Play()
	return s
}

func (s *TitleScene) Update() error {
	s.menu.Update()
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	ui.DrawDialogBox(
		screen,
		s.ctx.Assets.DialogBorder,
		menuX-8,
		menuY-8,
		208,
		72,
	)
	s.menu.Draw(screen, s.ctx.Assets.Font, menuX, menuY)
}

func (s *TitleScene) OnExit() {
	s.ctx.Assets.TitleBGM.Pause()
}
