package room

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/tilemap"
	"github.com/loneJogger/go-dungeon-crawler/internal/ui"
)

type Room struct {
	player    *entity.Player
	npcs      []*entity.NPC
	tileMap   *tiled.Map
	tilesets  []*ebiten.Image
	triggers  []tilemap.Trigger
	dialogBox *ui.DialogBox
	CameraX   float64
	CameraY   float64

	OnTrigger func(name string)
	OnEnter   func()
	OnExit    func()
}

func New(
	c *ctx.GameContext,
	player *entity.Player,
	npcs []*entity.NPC,
	tileMap *tiled.Map,
	tilesets []*ebiten.Image,
) *Room {
	return &Room{
		player:    player,
		npcs:      npcs,
		tileMap:   tileMap,
		tilesets:  tilesets,
		triggers:  tilemap.LoadTriggers(tileMap),
		dialogBox: ui.NewDialogBox(c.Assets.Font, c.Assets.DialogBorder),
	}
}

func (r *Room) ShowDialog(text string, onComplete func(), beep *audio.Player) {
	r.dialogBox.ShowText(text, onComplete, beep)
}

func (r *Room) SetPlayerPos(x, y float64) {
	r.player.X = x
	r.player.Y = y
}

func (r *Room) UpdateCamera() {
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

	r.CameraX = cx
	r.CameraY = cy
}

func (r *Room) Update() {
	if r.dialogBox.Active {
		r.dialogBox.Update()
		return
	}

	r.player.Update(r.tileMap, r.npcs)
	for _, npc := range r.npcs {
		npc.Update(r.tileMap, r.player)
	}
	r.UpdateCamera()

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		r.checkInteraction()
	}

	if r.OnTrigger != nil {
		trigger := tilemap.CheckTrigger(r.triggers, r.player.X+8, r.player.Y+12)
		if trigger != nil {
			r.OnTrigger(trigger.Name)
		}
	}
}

func (r *Room) Draw(screen *ebiten.Image) {
	tilemap.DrawMap(screen, r.tileMap, r.tilesets, r.CameraX, r.CameraY)
	for _, npc := range r.npcs {
		npc.OffsetX, npc.OffsetY = r.CameraX, r.CameraY
		npc.Draw(screen)
	}
	r.player.OffsetX, r.player.OffsetY = r.CameraX, r.CameraY
	r.player.Draw(screen)
	r.dialogBox.Draw(screen)
}

func (r *Room) checkInteraction() {
	checkX := r.player.X + config.TileSize/2
	checkY := r.player.Y + config.TileSize/2

	switch r.player.Direction {
	case 0:
		checkY += config.TileSize
	case 2:
		checkY -= config.TileSize
	case 1:
		if r.player.FacingRight {
			checkX += config.TileSize
		} else {
			checkX -= config.TileSize
		}
	}

	pt := image.Pt(int(checkX), int(checkY))
	for _, npc := range r.npcs {
		if pt.In(npc.Bounds()) && npc.OnInteract != nil {
			npc.OnInteract()
			return
		}
	}
}
