package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/state"
)

type MovementSystem struct{}

func (g MovementSystem) Update(w *ecs.World) {
	if w.State() != state.Play {
		return
	}

	for _, e := range w.Entities() {
		if !e.HasComponent(component.TransformType) || !e.HasComponent(component.VelocityType) {
			continue
		}

		v := e.GetComponent(component.VelocityType).(*component.Velocity)
		t := e.GetComponent(component.TransformType).(*component.Transform)

		t.X = t.X + int(v.X)
		t.Y = t.Y + int(v.Y)
	}
}
