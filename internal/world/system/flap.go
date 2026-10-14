package system

import (
	"log"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/input"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/state"
)

const FlapEvent ecs.EventType = "FLAPPINGTHEBIRD"

type FlapSystem struct{}

func HandleFlap(w *ecs.World, e ecs.Event) {
	for _, ent := range w.Entities() {
		f, ok := ecs.TryGet[*component.Flapper](ent)
		if !ok {
			continue
		}
		v, ok := ecs.TryGet[*component.Velocity](ent)
		if !ok {
			continue
		}
		v.Y = -f.Impulse
	}
}

// HandleStart begins a round, from either the title screen or the end screen.
func HandleStart(w *ecs.World, e ecs.Event) {
	w.SetNextState(state.Play)
}

func (f FlapSystem) Update(w *ecs.World) {
	if w.Resource(input.StateType).(*input.State).JustPressed(input.Flap) {
		log.Print("Flapping")
		w.Emit(ecs.Event{Verb: FlapEvent})
	}
}
