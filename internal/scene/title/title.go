package title

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/explore"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

const menuX = 100
const menuY = 80

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
			ss.SetScene(explore.New(ss, a))
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
		120,
		72,
	)

	// Draw Menu with Debug Print for now
	for i, item := range s.menu.Items {
		label := "  " + item.Label
		if i == s.menu.Focused() {
			label = "> " + item.Label
		}
		ebitenutil.DebugPrintAt(screen, label, menuX, menuY+i*20)
	}
}

func (s *TitleScene) OnExit() {
	s.assets.TitleBGM.Pause()
}
