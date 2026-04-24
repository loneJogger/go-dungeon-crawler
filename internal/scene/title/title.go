package title

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/location"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

const menuX = 64
const menuY = 144

type TitleScene struct {
	sceneSwitcher scene.SceneSwitcher
	menu          ui.Menu
	assets        *assets.Assets
}

func New(ss scene.SceneSwitcher, a *assets.Assets) *TitleScene {
	s := &TitleScene{sceneSwitcher: ss, assets: a}
	s.menu = *ui.NewMenu([]ui.MenuItem{
		{Label: "New Game", OnSelect: func() {
			a.GameStart.Rewind()
			a.GameStart.Play()
			ss.SetScene(location.NewTown(ss, a))
		}},
		{Label: "Continue", OnSelect: func() { /* TODO */ }},
		{Label: "Exit", OnSelect: func() { os.Exit(0) }},
	})
	s.menu.NavSound = a.MenuNav
	s.menu.SelectSound = a.MenuSelect
	a.TitleBGM.Play()
	return s
}

func (s *TitleScene) Update() error {
	s.menu.Update()
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {

	// Draw Dialog Box around Menu
	ui.DrawDialogBox(
		screen,
		s.assets.DialogBorder,
		menuX-8,
		menuY-8,
		208,
		72,
	)

	s.menu.Draw(screen, s.assets.Font, menuX, menuY)
}

func (s *TitleScene) OnExit() {
	s.assets.TitleBGM.Pause()
}
