package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)


func HandlePoison(w *ecs.World, event ecs.Event) { 
	if event.Verb != CollisionEvent {
		return
	}

	if hasPoisoned(event.Subject, event.Object) {
		w.Emit(ecs.Event{
			Verb: GameLossEvent,
		})
	}
}


func hasPoisoned(e, other *ecs.Entity) bool {
	if e.HasComponent(component.PoisoneeType) && other.HasComponent(component.PoisonerType) {
		return true
	}
	if e.HasComponent(component.PoisonerType) && other.HasComponent(component.PoisoneeType) {
		return true
	}
	return false
}
