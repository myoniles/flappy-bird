package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const FlapEvent ecs.EventType = "FLAPPINGTHEBIRD"

func init() { Handlers[FlapEvent] = append(Handlers[FlapEvent], HandleFlap) }

func HandleFlap(entities *[]*ecs.Entity, e ecs.Event) {
	for _, ent := range *entities {
		if !ent.HasComponent("Flapper") || !ent.HasComponent("Velocity") {
			continue
		}
		f := ent.GetComponent("Flapper").(*component.Flapper)
		v := ent.GetComponent("Velocity").(*component.Velocity)
		v.Y = -f.Impulse
	}
}
