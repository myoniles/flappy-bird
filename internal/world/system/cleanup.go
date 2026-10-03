package system

import (
	"slices"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

var OutOfBoundsEvent ecs.EventType = "OutOfBoundsEvent"

func init() {
	Handlers[OutOfBoundsEvent] = append(Handlers[OutOfBoundsEvent], HandleOutOfBounds)
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
		if !e.HasComponent("Transform") {
			continue
		}
		t := e.GetComponent("Transform").(*component.Transform)
		if isOutOfBounds(t.X, t.Y) {
			event := ecs.Event{Subject: e, Verb: OutOfBoundsEvent}
			q.Enqueue(event)
		}
	}
}

func isOutOfBounds(x, y int) bool {
	windowConf := config.GetWindowConfig()
	// keep it stupid simple, just go off of an arbitrary position
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
