package battle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
