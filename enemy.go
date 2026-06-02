package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X, Y  float64
	speed float32
}

func newEnemy() Enemy {
	return Enemy{
		X: float64(rand.Intn(screenW)),
		Y: float64(rand.Intn(100)),
		speed: 2,
	}
}

func (e *Enemy) Update() {
	e.Y += float64(e.speed)

	if e.Y > float64(screenH) {
		e.Y = 0
		e.X = float64(rand.Intn(screenW))
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	scale := 0.2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(enemyImg, op)
}