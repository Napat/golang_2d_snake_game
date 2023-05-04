package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGTH = 240
	TILE_SIZE     = 5
	SPEED_INIT    = 10
	SPEED_MAX     = 2
)

func main() {
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		gameOver: false,
		speed:    SPEED_INIT,
	}
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGTH*2) // x2 scale windows for display comfortably
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
