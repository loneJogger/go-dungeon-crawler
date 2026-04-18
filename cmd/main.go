package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
	"github.com/loneJogger/go-dungeon-crawler/internal/scene/explore"
)

func main() {
	g := game.New()
	g.SetScene(explore.New())

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("RPG Dungeon Crawler")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
