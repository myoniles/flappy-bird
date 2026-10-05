package system

import (
	"slices"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

var DestroyEvent ecs.EventType = "Destroy"

func init() {
	Handlers[DestroyEvent] = append(Handlers[DestroyEvent], HandleOutOfBounds)
}

func HandleOutOfBounds(entities *[]*ecs.Entity, e ecs.Event) {
	*entities = slices.DeleteFunc(*entities, func(x *ecs.Entity) bool {
		return x == e.Subject
	})
}

// TODO move to config
const OOBMod = 250

type CleanupSystem struct{}

func (c CleanupSystem) Update(entities []*ecs.Entity, q *ecs.EventQueue) {
	for _, e := range entities {
		if !e.HasComponent(component.TransformType) {
			continue
		}
		t := e.GetComponent(component.TransformType).(*component.Transform)
		if isOutOfBounds(t.X, t.Y) {
			event := ecs.Event{Subject: e, Verb: DestroyEvent}
			q.Enqueue(event)
		}
	}
}

func isOutOfBounds(x, y int) bool {
	windowConf := config.GetWindowConfig()
	// keep it stupid simple, just go off of an arbitrary position
	// but in all honesty I want the complexity
	// TODO MORE COMPLEX FOR NO REASON! USE SPRITE BOUNDING BOX FOR EFFICIENCY!!
	if x > windowConf.LogicalWidth+OOBMod {
		return true
	}

	if y > windowConf.LogicalHeight+OOBMod {
		return true
	}

	if x < -OOBMod {
		return true
	}

	if y < -OOBMod {
		return true
	}

	return false
}
