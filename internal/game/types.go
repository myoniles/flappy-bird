package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/render"
	"github.com/myoniles/vulcan/internal/world"
)

type Game struct {
	World world.World
}

func NewGame() *Game {
	g := Game{}
	g.World = world.NewWorld()
	return &g
}

func (g *Game) Update() error {
	g.World.Update(poll())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{63, 63, 225, 50})
	render.RenderWorld(g.World, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	cfg := config.GetWindowConfig()
	return cfg.LogicalWidth, cfg.LogicalHeight
}
