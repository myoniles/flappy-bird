package component

import "github.com/myoniles/vulcan/internal/ecs"

type Velocity struct {
	X, Y float64
}

func (v Velocity) Type() ecs.ComponentType {
	return "Velocity"
}
