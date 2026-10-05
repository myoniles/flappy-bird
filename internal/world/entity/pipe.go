package entity

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand/v2"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

// TODO move to config
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
	component.Transform
}

func NewPipePair() []*ecs.Entity {
	windowConf := config.GetWindowConfig()
	offset := rand.IntN(3 * windowConf.LogicalHeight / 4)
	topImg, bottomImg := bodyPairs(offset)

	bottomHeight := bottomImg.Bounds().Dy()
	xIndex := windowConf.LogicalWidth + PipeWidth/2

	top := ecs.NewEntity(PipePrefix)
	top.AddComponent(&component.Transform{X: xIndex, Y: offset / 2})
	top.AddComponent(component.NewStaticSprite(&topImg))
	top.AddComponent(&component.Velocity{X: -5.0, Y: 0})
	top.AddComponent(&component.Poisoner{})
	top.AddComponent(&component.Collider{Width: PipeWidth, Height: offset, Layer: component.LayerObstacle})

	bottom := ecs.NewEntity(PipePrefix)
	bottom.AddComponent(&component.Transform{X: xIndex, Y: windowConf.LogicalHeight - bottomHeight/2})
	bottom.AddComponent(component.NewStaticSprite(&bottomImg))
	bottom.AddComponent(&component.Velocity{X: -5.0, Y: 0})
	bottom.AddComponent(&component.Poisoner{})
	top.AddComponent(&component.Collider{Width: PipeWidth, Height: bottomHeight, Layer: component.LayerObstacle})

	return []*ecs.Entity{top, bottom}
}

func bodyPairs(offset int) (image.Image, image.Image) {
	windowConf := config.GetWindowConfig()
	topImg := pipeBody(offset)
	bottomImg := pipeBody(windowConf.LogicalHeight - (offset + GapHeight))
	return topImg, bottomImg
}

func pipeBody(height int) image.Image {
	r := image.Rect(0, 0, PipeWidth, height)
	rgba := image.NewRGBA(r)
	var green color.Color
	green = color.RGBA{0, 255, 0, 255}
	greenSource := image.NewUniform(green)
	// This writes to rgba
	draw.Draw(rgba, r.Bounds(), greenSource, image.Point{}, draw.Src)
	return rgba
}
