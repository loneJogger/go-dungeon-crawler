package title

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/explore"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

type TitleScene struct {
	sceneSwitcher scene.SceneSwitcher
	menu          ui.Menu
}

func New(ss scene.SceneSwitcher) *TitleScene {
	s := &TitleScene{sceneSwitcher: ss}
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

	// Draw with Debug Print for now
	for i, item := range s.menu.Items {
		label := "  " + item.Label
		if i == s.menu.Focused() {
			label = "> " + item.Label
		}
		ebitenutil.DebugPrintAt(screen, label, 100, 80+i*20)
	}
}
