package component

import "github.com/myoniles/vulcan/internal/world"

type Position struct {
	X, Y     int
	rotation float32
	scale    float32
}

func (p Position) Type() world.ComponentType {
	return "Transform"
}
