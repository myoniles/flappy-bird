package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type GravitySystem struct{}

func (g GravitySystem) Update(w ecs.World) {
	if w.State() != ecs.PlayState {
		return
	}
	for _, e := range w.Entities() {
		if !e.HasComponent(component.TransformType) || !e.HasComponent(component.VelocityType) || !e.HasComponent(component.GravityType) {
			continue
		}

		v := e.GetComponent(component.VelocityType).(*component.Velocity)
		grav := e.GetComponent(component.GravityType).(component.Gravity)

		v.Y += grav.Accel * (1.0/60.0)
	}
}
