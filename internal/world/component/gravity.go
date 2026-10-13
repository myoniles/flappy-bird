package component

import "github.com/myoniles/vulcan/internal/ecs"

const GravityType ecs.ComponentType = "Gravity"

type Gravity struct {
	Accel float64
}

func (g *Gravity) Type() ecs.ComponentType {
	return GravityType
}
