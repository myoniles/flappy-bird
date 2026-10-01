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

var img image.Image

func NewBird() *ecs.Entity {
	e := ecs.NewEntity("bird")
	e.AddComponent(component.NewStaticSprite(loadBirdImage()))
	e.AddComponent(component.Position{X: 0, Y: 0})
	return e
}

func loadBirdImage() *image.Image {
	if img == nil {
		log.Print("Loading Bird image")
		i, _, err := image.Decode(bytes.NewReader(flappy.Gopher_png))
		if err != nil {
			log.Print(err)
			log.Fatal("Unabled to load the gopher image. Load bearing png")
		}
		img = i
	}
	return &img
}
