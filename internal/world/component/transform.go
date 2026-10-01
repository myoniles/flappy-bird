package component

import "github.com/myoniles/vulcan/internal/ecs"

type Position struct {
	X, Y     int
	rotation float32
	scale    float32
}

func (p Position) Type() ecs.ComponentType {
	return "Transform"
}
