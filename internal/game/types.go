package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Mode GameMode
}


func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{63, 63, 225, 50})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
