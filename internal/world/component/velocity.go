package component

import "github.com/myoniles/vulcan/internal/ecs"

const VelocityType ecs.ComponentType = "Velocity"

type Velocity struct {
	X, Y float64
}

func (v Velocity) Type() ecs.ComponentType {
	return VelocityType
}
