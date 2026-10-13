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
	if w.State() != state.Play {
		return
	}
	for _, ent := range w.Entities() {
		if !ent.HasComponent(component.FlapperType) || !ent.HasComponent(component.VelocityType) {
			continue
		}
		f := ent.GetComponent(component.FlapperType).(*component.Flapper)
		v := ent.GetComponent(component.VelocityType).(*component.Velocity)
		v.Y = -f.Impulse
	}
}

func HandleTitleFlap(w *ecs.World, e ecs.Event) {
	if w.State() != state.Title {
		return
	}
	w.SetNextState(state.Play)
}

func (f FlapSystem) Update(w *ecs.World) {
	if w.Resource(input.StateType).(*input.State).JustPressed(input.Flap) {
		log.Print("Flapping")
		w.Emit(ecs.Event{Verb: FlapEvent})
	}
}

func HandleEndFlap(w *ecs.World, e ecs.Event) {
	log.Print("Handling Restart")
	if w.State() != state.End {
		return
	}
	w.SetNextState(state.Play)
}
