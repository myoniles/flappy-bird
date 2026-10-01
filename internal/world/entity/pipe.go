package entity

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand/v2"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const (
	PipePrefix = "pipe"
	PipeWidth  = 50
	GapHeight  = 150
)

var PipeGap = config.GetWindowConfig().LogicalWidth / 4

type Pipe struct {
	Top    halfPipe
	Bottom halfPipe
}

type halfPipe struct {
	component.Sprite
	component.Position
}

func NewPipePair() []*ecs.Entity {
	windowConf := config.GetWindowConfig()
	offset := rand.IntN(3 * windowConf.LogicalHeight / 4)
	topImg, bottomImg := bodyPairs(offset)

	bottomHeight := bottomImg.Bounds().Dy()

	top := ecs.NewEntity(PipePrefix)
	top.AddComponent(component.Position{X: windowConf.LogicalWidth, Y: offset / 2})
	top.AddComponent(component.NewStaticSprite(&topImg))

	bottom := ecs.NewEntity(PipePrefix)
	bottom.AddComponent(component.Position{X: windowConf.LogicalWidth, Y: windowConf.LogicalHeight - bottomHeight/2})
	bottom.AddComponent(component.NewStaticSprite(&bottomImg))

	return []*ecs.Entity{top, bottom}
}

func bodyPairs(offset int) (image.Image, image.Image) {
	windowConf := config.GetWindowConfig()
	topImg := pipeBody(offset, true)
	bottomImg := pipeBody(windowConf.LogicalHeight - (offset + GapHeight), false)
	return topImg, bottomImg
}

func pipeBody(height int, top bool) (image.Image) {
	log.Print("Rendering Pipe image\n")
	r := image.Rect(0, 0, PipeWidth, height)
	rgba := image.NewRGBA(r)
	var green color.Color
	if top {
		green = color.RGBA{0, 255, 0, 255}
	} else {
		green = color.RGBA{255, 0, 0, 255}
	}
	greenSource := image.NewUniform(green)
	// This writes to rgba
	draw.Draw(rgba, r.Bounds(), greenSource, image.Point{}, draw.Src)
	return rgba
}
