package room

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
)

type Room struct {
	player   *entity.Player
	tileMap  *tiled.Map
	tilesets []*ebiten.Image
	cameraX  float64
	cameraY  float64
}

func (r *Room) updateCamera() {
	mapW := float64(r.tileMap.Width * r.tileMap.TileWidth)
	mapH := float64(r.tileMap.Height * r.tileMap.TileHeight)

	cx := r.player.X - config.ScreenWidth/2 + config.TileSize/2
	cy := r.player.Y - config.ScreenHeight/2 + config.TileSize/2

	if cx < 0 {
		cx = 0
	} else if cx > mapW-config.ScreenWidth {
		cx = mapW - config.ScreenWidth
	}
	if cy < 0 {
		cy = 0
	} else if cy > mapH-config.ScreenHeight {
		cy = mapH - config.ScreenHeight
	}

	r.cameraX = cx
	r.cameraY = cy
}
