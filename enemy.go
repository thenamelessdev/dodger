package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (e *Enemy) Update(g *Game) {
	e.Y += float64(e.speed)

	if e.Y > float64(screenH) {
		e.Y = 0
		e.X = float64(rand.Intn(screenW))
	}

	x1, y1, w1, h1 := e.Bounds()
	x2, y2, w2, h2 := g.player.Bounds()
	if Overlaps(x1, y1, w1, h1, x2, y2, w2, h2) {
		fmt.Println("Collision ", rand.Int())
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	scale := 0.2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(enemyImg, op)

	if os.Getenv("DEBUG") != "" {
		x, y, w, h := e.Bounds()
		vector.StrokeRect(screen, float32(x), float32(y), float32(w), float32(h), 1, color.RGBA{255, 0, 0, 255}, false)
	}
}

func (e *Enemy) Bounds() (x, y, w, h float64) {
    w = float64(enemyImg.Bounds().Dx()) * 0.2
    h = float64(enemyImg.Bounds().Dy()) * 0.2
    return e.X, e.Y, w, h
}