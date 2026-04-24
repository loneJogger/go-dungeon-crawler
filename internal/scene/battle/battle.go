package battle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

type BattleScene struct{}

func New() *BattleScene {
	return &BattleScene{}
}

func (s *BattleScene) Update() error {
	return nil
}

func (s *BattleScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Battle Scene")
}

func (s *BattleScene) TransitionPhase() transition.Phase {
	return transition.Closing
}

func (s *BattleScene) TransitionType() game.TransitionType {
	return game.TransitionSpiral
}
