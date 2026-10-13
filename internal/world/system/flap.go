package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/input"
	"github.com/myoniles/vulcan/internal/world/component"
)

const FlapEvent ecs.EventType = "FLAPPINGTHEBIRD"

type FlapSystem struct{}

func init() {
	Handlers[FlapEvent] = append(Handlers[FlapEvent], HandleFlap)
	Handlers[FlapEvent] = append(Handlers[FlapEvent], HandleTitleFlap)
}

func HandleFlap(w ecs.World, e ecs.Event) {
	if w.State() != ecs.PlayState {
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

func HandleTitleFlap(w ecs.World, e ecs.Event) {
	if w.State() != ecs.Title {
		return
	}
	w.SetNextState(ecs.PlayState)
}

func (f FlapSystem) Update(w ecs.World) {
	if w.GetInput().JustPressed(input.Flap) {
		w.Emit(ecs.Event{Verb: FlapEvent})
	}
}
