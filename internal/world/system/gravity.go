package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type GravitySystem struct{}

func (g GravitySystem) Update(entries []*ecs.Entity) {
	for _, e := range entries {
		if !e.HasComponent("Transform") || !e.HasComponent("Velocity") || !e.HasComponent("Gravity") {
			continue
		}

		v := e.GetComponent("Velocity").(*component.Velocity)
		grav := e.GetComponent("Gravity").(component.Gravity)

		// Ebiten ticks at a fixed 60FPS
		// TODO make this configurable
		v.Y += grav.Accel * (1.0/60.0)
	}
}
