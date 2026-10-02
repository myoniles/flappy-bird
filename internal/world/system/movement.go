package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type MovementSystem struct{}

func (g MovementSystem) Update(entries []*ecs.Entity, q *ecs.EventQueue) {
	for _, e := range entries {
		if !e.HasComponent("Transform") || !e.HasComponent("Velocity") {
			continue
		}

		v := e.GetComponent("Velocity").(*component.Velocity)
		t := e.GetComponent("Transform").(*component.Transform)

		t.X = t.X + int(v.X)
		t.Y = t.Y + int(v.Y)
	}
}
