package ctx

import (
	"github.com/loneJogger/go-dungeon-crawler/internal/assets"
	"github.com/loneJogger/go-dungeon-crawler/internal/party"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene"
)

type GameContext struct {
	Assets *assets.Assets
	Party  *party.Party
	SS     scene.SceneSwitcher
}
