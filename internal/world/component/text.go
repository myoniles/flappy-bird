package component

import (
	"image/color"

	"github.com/myoniles/vulcan/internal/ecs"
)

const TextType ecs.ComponentType = "TextType"

type Text struct {
	FontSize int
	Text     string
	Color    color.Color
}

func (t *Text) Type() ecs.ComponentType {
	return TextType
}
