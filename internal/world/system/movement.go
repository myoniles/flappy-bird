package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type MovementSystem struct{}

func (g MovementSystem) Update(w *ecs.World) {
	for _, e := range w.Entities() {
		v, ok := ecs.TryGet[*component.Velocity](e)
		if !ok {
			continue
		}
		t, ok := ecs.TryGet[*component.Transform](e)
		if !ok {
			continue
		}

		t.X = t.X + int(v.X)
		t.Y = t.Y + int(v.Y)
	}
}
