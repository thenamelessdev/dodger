package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenW = 640
	screenH = 360
)

type Game struct {
	gameOver bool;
	player Player;
}

func newGame() *Game {
	return &Game{
		gameOver: false,
		player: newPlayer(10),
	}
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	g.player.Draw(screen)
}

func (g *Game) Layout(ow, oh int) (int, int) { return ow, oh }