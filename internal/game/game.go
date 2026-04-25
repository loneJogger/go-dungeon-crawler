package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/party"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/title"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type TransitionStarter interface {
	TransitionPhase() transition.Phase
	TransitionType() transition.TransitionType
}

type Game struct {
	current    scene.Scene
	pending    scene.Scene
	transition transition.Transition
	ctx        *ctx.GameContext
}

func New() (*Game, error) {
	a, err := assets.LoadAssets()
	if err != nil {
		return nil, err
	}

	g := &Game{}
	g.ctx = &ctx.GameContext{
		Assets: a,
		Party:  party.New(),
		SS:     g,
	}
	g.SetScene(title.New(g.ctx))
	return g, nil
}

func (g *Game) SetScene(s scene.Scene) {
	if g.current == nil {
		g.current = s
		return
	}
	if g.transition != nil {
		return
	}
	if ex, ok := g.current.(scene.SceneExiter); ok {
		ex.OnExit()
	}
	g.pending = s

	phase := transition.Closing
	tType := transition.TransitionSpiral
	if ts, ok := s.(TransitionStarter); ok {
		phase = ts.TransitionPhase()
		tType = ts.TransitionType()
	}

	switch tType {
	case transition.TransitionBox:
		g.transition = transition.NewBox(phase)
	default:
		g.transition = transition.New(phase)
	}
}

func (g *Game) Update() error {
	if g.transition != nil {
		g.transition.Update()
		if g.transition.IsFullyBlack() && g.pending != nil {
			g.current = g.pending
			g.pending = nil
		}
		if g.transition.IsDone() {
			g.transition = nil
			if en, ok := g.current.(scene.SceneEnter); ok {
				en.OnEnter()
			}
		}
		return nil
	}
	return g.current.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current.Draw(screen)
	if g.transition != nil {
		g.transition.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}
