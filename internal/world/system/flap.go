package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const FlapEvent ecs.EventType = "FLAPPINGTHEBIRD"

func init() { Handlers[FlapEvent] = append(Handlers[FlapEvent], HandleFlap) }

func HandleFlap(w ecs.World, e ecs.Event) {
	for _, ent := range w.Entities() {
		if !ent.HasComponent(component.FlapperType) || !ent.HasComponent(component.VelocityType) {
			continue
		}
		f := ent.GetComponent(component.FlapperType).(*component.Flapper)
		v := ent.GetComponent(component.VelocityType).(*component.Velocity)
		v.Y = -f.Impulse
	}
}
