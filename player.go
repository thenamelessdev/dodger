package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	x float64
}

func newPlayer(x float64) Player {
	return Player{x: x}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {p.x -= 1}
	if ebiten.IsKeyPressed(ebiten.KeyD) {p.x += 1}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		p.x,
		100,
	)
	screen.DrawImage(playerImg, op)
}