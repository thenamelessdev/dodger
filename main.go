package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Dodger")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}