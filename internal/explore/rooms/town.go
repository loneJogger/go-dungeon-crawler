package rooms

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/loneJogger/go-dungeon-crawler/internal/config"
	"github.com/loneJogger/go-dungeon-crawler/internal/ctx"
	"github.com/loneJogger/go-dungeon-crawler/internal/entity"
	"github.com/loneJogger/go-dungeon-crawler/internal/explore/room"
	"github.com/loneJogger/go-dungeon-crawler/internal/transition"
)

const TownStartX = 96
const TownStartY = 200

func NewTownRoom(c *ctx.GameContext, switchRoom SwitchFn) *room.Room {
	p := entity.NewPlayer(0, 0, c.Assets.PCSprite)

	thief := entity.NewNPC(128, 160, 0)
	thief.Image = c.Assets.NPCThief
	thief.Wanders = true

	blackBelt := entity.NewNPC(272, 176, 1)
	blackBelt.Image = c.Assets.NPCBlackBelt
	blackBelt.Wanders = false

	r := room.New(c, p, []*entity.NPC{thief, blackBelt}, c.Assets.TownMap, []*ebiten.Image{c.Assets.TownTileset})

	thief.OnInteract = func() {
		r.ShowDialog(
			"Hey You!\n\nThere are some {COLOR PINK}strange characters{COLOR DEFAULT} at the inn...\n\nuh...\nMaybe go do something about it?",
			nil,
			c.Assets.VoiceOne,
		)
	}

	blackBelt.OnInteract = func() {
		r.ShowDialog(
			"... ... ... ... ... ... ... ...\n\nIf I was a {COLOR ICY_BLUE}sea gull{COLOR DEFAULT}, I would {COLOR MINT}fly{COLOR DEFAULT} as far as I could!\n\nI would fly to {COLOR PINK}far away places{COLOR DEFAULT} and {COLOR BLOODY}sing{COLOR DEFAULT} for many people!\n\n",
			nil,
			c.Assets.VoiceTwo,
		)
	}

	r.OnTrigger = func(name string) {
		switch name {
		case "inn_entrance":
			retX, retY := p.X, p.Y+config.TileSize
			switchRoom(NewInnRoom(c, switchRoom, retX, retY), transition.TransitionBox, 240, 176)
		case "west_shop_entrance":
			switchRoom(NewShopRoom(c, switchRoom, 144, 144, 224, 144), transition.TransitionBox, 128, 176)
		case "east_shop_entrance":
			switchRoom(NewShopRoom(c, switchRoom, 144, 144, 224, 144), transition.TransitionBox, 528, 176)
		}
	}

	return r
}
