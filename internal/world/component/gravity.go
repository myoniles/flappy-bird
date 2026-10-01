package component

import "github.com/myoniles/vulcan/internal/ecs"

type Gravity struct {
	Accel float64
}

func (g Gravity) Type() ecs.ComponentType {
	return "Gravity"
}

