package component

import "github.com/myoniles/vulcan/internal/ecs"

const TransformType ecs.ComponentType = "Transform"

type Transform struct {
	X, Y     int
	rotation float32
	scale    float32
}

func (p Transform) Type() ecs.ComponentType {
	return TransformType
}
