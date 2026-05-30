package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var playerImg *ebiten.Image

func init() {
	playerImg = loadImage("assets/Player.png")
}

func loadImage(path string) *ebiten.Image {
    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        log.Fatal(err)
    }
    return ebiten.NewImageFromImage(img)
}