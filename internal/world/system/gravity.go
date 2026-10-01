package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type GravitySystem struct{}

func (g *GravitySystem) Update(entries []*ecs.Entity, dt float64) {
	for _, e := range entries {
		if !e.HasComponent("Transform") || !e.HasComponent("Velocity") || !e.HasComponent("Gravity") {
			continue
		}

		v := e.GetComponent("Velocity").(*component.Velocity)
		grav := e.GetComponent("Gravity").(component.Gravity)

		v.Y += grav.Accel * dt
	}
}
