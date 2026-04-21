package mapscene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
	"github.com/loneJogger/go-dungeon-crawler/internal/world"
)

const (
	TileSize = 16
)

var PinkText = []float32{0.98, 0.653, 0.724, 1}

type MapScene struct {
	sceneSwitcher scene.SceneSwitcher
	assets        *assets.Assets
	player        *entity.Player
	npcs          []*entity.NPC
	dialogBox     *ui.DialogBox
	triggers      []world.Trigger
	tileMap       *tiled.Map
	tileset       *ebiten.Image
}

func NewMapScene(
	ss scene.SceneSwitcher,
	a *assets.Assets,
	p *entity.Player,
	npcs []*entity.NPC,
	tm *tiled.Map,
) *MapScene {
	ms := &MapScene{
		sceneSwitcher: ss,
		assets:        a,
		player:        p,
		npcs:          npcs,
		tileMap:       tm,
	}
	ms.player.Update(ms.tileMap, ms.npcs)
	return ms
}
