package system

import (
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const DestroyEvent ecs.EventType = "Destroy"

func HandleOutOfBounds(w *ecs.World, e ecs.Event) {
	w.RemoveEntity(e.Subject)
}

type CleanupSystem struct{}

func (c CleanupSystem) Update(w *ecs.World) {
	entities := w.Query(component.TransformType)
	for _, e := range entities {
		t, ok := ecs.Get[*component.Transform](e)
		if !ok {
			continue
		}
		if isOutOfBounds(t.X, t.Y) {
			event := ecs.Event{Subject: e, Verb: DestroyEvent}
			w.Emit(event)
		}
	}

	entities = w.Query(component.StateScopedType)
	for _, e := range entities {
		scope, ok := ecs.Get[*component.StateScoped](e)
		if !ok {
			continue
		}
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
	if x > windowConf.LogicalWidth+config.OutOfBounds {
		return true
	}

	if y > windowConf.LogicalHeight+config.OutOfBounds {
		return true
	}

	if x < -config.OutOfBounds {
		return true
	}

	if y < -config.OutOfBounds {
		return true
	}

	return false
}
