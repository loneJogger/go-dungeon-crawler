package title

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		{Label: "New Game", OnSelect: func() { ss.SetScene(explore.New()) }},
		{Label: "Contine", OnSelect: func() { /* TODO */ }},
		{Label: "Exit", OnSelect: func() { os.Exit(0) }},
	})
	return s
}

func (s *TitleScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		s.menu.MoveUp()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		s.menu.MoveDown()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		s.menu.Select()
	}
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
