package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	x float64
}

func newPlayer(x float64) Player {
	return Player{x: x}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) && p.x > 0 {p.x -= 4}
	if ebiten.IsKeyPressed(ebiten.KeyD) && p.x < float64(screenW) {p.x += 4}
}

func (p *Player) Reset() {
	p.x = 10
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(
		p.x,
		300,
	)
	screen.DrawImage(playerImg, op)

	if os.Getenv("DEBUG") != "" {
		x, y, w, h := p.Bounds()
		vector.StrokeRect(screen, float32(x), float32(y), float32(w), float32(h), 1, color.RGBA{255, 0, 0, 255}, false)
	}
}

func (p *Player) Bounds() (x, y, w, h float64) {
    w = float64(playerImg.Bounds().Dx()) * 0.1
    h = float64(playerImg.Bounds().Dy()) * 0.1
    return p.x, 300, w, h
}