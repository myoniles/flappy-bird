package entity

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/images/flappy"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func NewBird() *ecs.Entity {
	e := ecs.NewEntity()
	img, _, err := image.Decode(bytes.NewReader(flappy.Gopher_png))
	if err != nil {
		log.Print(err)
		log.Fatal("Unabled to load the gopher image. Load bearing png")
	}
	e.AddComponent(component.NewStaticSprite(&img))
	e.AddComponent(component.Position{X: 0, Y: 0})
	return e
}
