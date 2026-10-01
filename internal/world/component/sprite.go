package component

import "github.com/myoniles/vulcan/internal/world"

// We are only using static sprites, I am sure this is more complicated than what I am doing
type Sprite struct {
	img []byte
}

func NewStaticSprite(s []byte) Sprite {
	return Sprite{s}
}

func (s Sprite) Type() world.ComponentType {
	return "sprite"
}
