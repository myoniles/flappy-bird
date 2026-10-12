package system

import (
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const DestroyEvent ecs.EventType = "Destroy"

func init() {
	Handlers[DestroyEvent] = append(Handlers[DestroyEvent], HandleOutOfBounds)
}

func HandleOutOfBounds(w ecs.World, e ecs.Event) {
	w.RemoveEntity(e.Subject)
}

// TODO move to config
const OOBMod = 250

type CleanupSystem struct{}

func (c CleanupSystem) Update(w ecs.World) {
	entities := w.Query(component.TransformType)
	for _, e := range entities {
		t := e.GetComponent(component.TransformType).(*component.Transform)
		if isOutOfBounds(t.X, t.Y) {
			event := ecs.Event{Subject: e, Verb: DestroyEvent}
			w.Emit(event)
		}
	}

	entities = w.Query(component.StateScopedType)
	for _, e := range entities {
		scope := e.GetComponent(component.StateScopedType).(*component.StateScoped)
		if w.State() != scope.Active {
			event := ecs.Event{Subject: e, Verb: DestroyEvent}
			w.Emit(event)
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
