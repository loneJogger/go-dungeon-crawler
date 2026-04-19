package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type Game struct {
	current    scene.Scene
	pending    scene.Scene
	transition *transition.SpiralTransition
}

type TransitionStarter interface {
	TransitionPhase() transition.Phase
}

func New() *Game {
	return &Game{}
}

func (g *Game) SetScene(s scene.Scene) {
	if g.current == nil {
		g.current = s
		return
	}
	if ex, ok := g.current.(scene.SceneExiter); ok {
		ex.OnExit()
	}
	g.pending = s
	phase := transition.Closing
	if ts, ok := s.(TransitionStarter); ok {
		phase = ts.TransitionPhase()
	}
	g.transition = transition.New(phase)
}

func (g *Game) Update() error {
	if g.transition != nil {
		g.transition.Update()
		// when fully black, swap the scene
		if g.transition.IsFullyBlack() && g.pending != nil {
			g.current = g.pending
			g.pending = nil
		}
		if g.transition.Done {
			g.transition = nil
			if en, ok := g.current.(scene.SceneEnter); ok {
				en.OnEnter()
			}
		}
		return nil // block input during transition
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
