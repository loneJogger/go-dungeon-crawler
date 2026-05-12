package title

import (
	"image/color"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

const menuX = 64
const menuY = 144
const titleDelay = 520
const titleSolidAt = 1500
const flashDuration = 3

type titleState int

const (
	titleWaiting titleState = iota
	titleFlash
	titleFlicker
	titleSolid
)

type TitleScene struct {
	ctx  *ctx.GameContext
	menu ui.Menu

	state       titleState
	tick        int
	flashTick   int
	flickerOn   bool
	flickerTick int
	flickerDur  int
	graphicX    int
	graphicY    int
}

func New(c *ctx.GameContext) *TitleScene {
	s := &TitleScene{ctx: c}
	s.menu = *ui.NewMenu([]ui.MenuItem{
		{Label: "New Game", OnSelect: func() {
			c.Assets.GameStart.Rewind()
			c.Assets.GameStart.Play()
			c.ScSwitcher.SetScene(explore.New(c))
		}},
		{Label: "Continue", OnSelect: func() { /* TODO */ }},
		{Label: "Exit", OnSelect: func() { os.Exit(0) }},
	})
	s.menu.NavSound = c.Assets.MenuNav
	s.menu.SelectSound = c.Assets.MenuSelect

	gw := c.Assets.TitleGraphic.Bounds().Dx()
	gh := c.Assets.TitleGraphic.Bounds().Dy()
	s.graphicX = (config.ScreenWidth - gw) / 2
	s.graphicY = (config.ScreenHeight/2 - gh) / 2

	c.Assets.TitleBGM.Play()
	return s
}

func (s *TitleScene) Update() error {
	s.menu.Update()
	s.tick++

	switch s.state {
	case titleWaiting:
		if s.tick >= titleDelay {
			s.state = titleFlash
		}
	case titleFlash:
		s.flashTick++
		if s.flashTick >= flashDuration {
			s.state = titleFlicker
			s.flickerOn = true
			s.flickerDur = rand.Intn(2) + 1
		}
	case titleFlicker:
		if s.tick >= titleSolidAt {
			s.state = titleSolid
			break
		}
		s.flickerTick++
		if s.flickerTick >= s.flickerDur {
			s.flickerTick = 0
			s.flickerOn = !s.flickerOn
			if s.flickerOn {
				s.flickerDur = rand.Intn(2) + 1
			} else {
				s.flickerDur = rand.Intn(3) + 3
			}
		}
	}

	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	switch s.state {
	case titleFlash:
		screen.Fill(color.White)
	case titleFlicker:
		if s.flickerOn {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(s.graphicX), float64(s.graphicY))
			screen.DrawImage(s.ctx.Assets.TitleGraphic, op)
		}
	case titleSolid:
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(s.graphicX), float64(s.graphicY))
		screen.DrawImage(s.ctx.Assets.TitleGraphic, op)
	}

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
