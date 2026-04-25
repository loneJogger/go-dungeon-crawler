package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loneJogger/go-dungeon-crawler/internal/game"
)

func main() {
	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("RPG Dungeon Crawler")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
