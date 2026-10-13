package component

import (
	"image"

	"github.com/myoniles/vulcan/internal/ecs"
)

// We are only using static sprites, I am sure this is more complicated than what I am doing
const SpriteType ecs.ComponentType = "Sprite"

type Sprite struct {
	Img image.Image
}

func NewStaticSprite(i image.Image) *Sprite {
	return &Sprite{i}
}

func (s *Sprite) Type() ecs.ComponentType {
	return SpriteType
}
