package asset

import (
	"image"
	"image/color"
	"image/draw"
)

func NewRect(width, height int, c color.Color) image.Image {
	r := image.Rect(0, 0, width, height)
	rgba := image.NewRGBA(r)
	colorSource := image.NewUniform(c)
	draw.Draw(rgba, r.Bounds(), colorSource, image.Point{}, draw.Src)
	return rgba
}
