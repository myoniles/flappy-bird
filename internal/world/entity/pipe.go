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
	PipePrefix  = "pipe"
	ScorePrefix = "coin"
	PipeWidth   = 50
	GapHeight   = 150
)

var PipeGap = config.GetWindowConfig().LogicalWidth / 4

type Pipe struct {
	Top       halfPipe
	Bottom    halfPipe
	Invisible ecs.Entity
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

	v := &component.Velocity{X: -5, Y: 0}
	p := &component.Poisoner{}

	top := ecs.NewEntity(PipePrefix)
	top.AddComponent(v)
	top.AddComponent(p)
	top.AddComponent(&component.Transform{X: xIndex, Y: offset / 2})
	top.AddComponent(component.NewStaticSprite(&topImg))
	top.AddComponent(
		&component.Collider{
			Width:  PipeWidth,
			Height: offset,
			Layer:  component.LayerObstacle,
		})

	bottom := ecs.NewEntity(PipePrefix)
	bottom.AddComponent(v)
	bottom.AddComponent(p)
	bottom.AddComponent(
		&component.Transform{
			X: xIndex,
			Y: windowConf.LogicalHeight - bottomHeight/2,
		})
	bottom.AddComponent(component.NewStaticSprite(&bottomImg))
	bottom.AddComponent(
		&component.Collider{
			Width:  PipeWidth,
			Height: bottomHeight,
			Layer:  component.LayerObstacle,
		})

	// What actually scores the player a point
	inv := ecs.NewEntity(ScorePrefix)
	inv.AddComponent(&component.Collider{Width: PipeWidth, Height: GapHeight, Layer: component.LayerObstacle})
	inv.AddComponent(&component.Transform{X: xIndex, Y: offset + GapHeight/2})
	inv.AddComponent(&component.Consumable{})
	inv.AddComponent(v)
	img := pipeBody(GapHeight, true)
	inv.AddComponent(component.NewStaticSprite(&img))

	return []*ecs.Entity{top, bottom, inv}
}

func bodyPairs(offset int) (image.Image, image.Image) {
	windowConf := config.GetWindowConfig()
	topImg := pipeBody(offset, false)
	bottomImg := pipeBody(windowConf.LogicalHeight-(offset+GapHeight), false)
	return topImg, bottomImg
}

func pipeBody(height int, coin bool) image.Image {
	r := image.Rect(0, 0, PipeWidth, height)
	rgba := image.NewRGBA(r)
	var green color.Color
	if coin {
		green = color.RGBA{225, 0, 0, 255}
	} else {
		green = color.RGBA{0, 255, 0, 255}
	}
	greenSource := image.NewUniform(green)
	// This writes to rgba
	draw.Draw(rgba, r.Bounds(), greenSource, image.Point{}, draw.Src)
	return rgba
}
