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
	Enemies []Enemy
}

func newGame() *Game {
	g := &Game{
		gameOver: false,
		player: newPlayer(10),
	}

	for i := 0; i < 7; i++ {
		g.Enemies = append(g.Enemies, newEnemy())
	}

	return g
}

func (g *Game) Update() error {
	if g.gameOver {
		for e := range g.Enemies {
			g.Enemies[e].Reset()
		}
		g.player.Reset()
		g.gameOver = false
	}
	g.player.Update()
	for i := range g.Enemies {
		g.Enemies[i].Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	g.player.Draw(screen)
	for i := range g.Enemies {
		g.Enemies[i].Draw(screen)
	}
}

func (g *Game) Layout(ow, oh int) (int, int) { return ow, oh }