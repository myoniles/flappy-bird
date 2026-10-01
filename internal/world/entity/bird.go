package entity

import (
	"github.com/hajimehoshi/ebiten/examples/resources/images/flappy"
	"github.com/myoniles/vulcan/internal/world"
	"github.com/myoniles/vulcan/internal/world/component"
)

func NewBird() world.Entity {
	e := world.Entity{}

	e.AddComponent(component.NewStaticSprite(flappy.Gopher_png))
	e.AddComponent(component.Position{X: 0, Y:0})
	return e
}
