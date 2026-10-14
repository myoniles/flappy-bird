package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type GravitySystem struct{}

func (g GravitySystem) Update(w *ecs.World) {
	for _, e := range w.Entities() {
		if !e.HasComponent(component.TransformType) {
			continue
		}
		v, ok := ecs.TryGet[*component.Velocity](e)
		if !ok {
			continue
		}
		grav, ok := ecs.TryGet[*component.Gravity](e)
		if !ok {
			continue
		}

		v.Y += grav.Accel * (1.0 / 60.0)
	}
}
