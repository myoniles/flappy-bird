package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type GravitySystem struct{}

func (g GravitySystem) Update(entries []*ecs.Entity, q *ecs.EventQueue) {
	for _, e := range entries {
		if !e.HasComponent(component.TransformType) || !e.HasComponent(component.VelocityType) || !e.HasComponent(component.GravityType) {
			continue
		}

		v := e.GetComponent(component.VelocityType).(*component.Velocity)
		grav := e.GetComponent(component.GravityType).(component.Gravity)

		// Ebiten ticks at a fixed 60FPS
		// TODO make this configurable
		v.Y += grav.Accel * (1.0/60.0)
	}
}
