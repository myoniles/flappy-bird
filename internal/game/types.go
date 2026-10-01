package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/myoniles/vulcan/internal/render"
	"github.com/myoniles/vulcan/internal/world"
)

type Game struct {
	Mode GameMode
	World world.World
}

func NewGame() *Game {
	g := Game{}
	g.Mode = GameStart
	g.World = world.NewWorld()
	return &g
}


func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{63, 63, 225, 50})
	render.RenderWorld(g.World, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
