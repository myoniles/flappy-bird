package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const SpawnEvent ecs.EventType = "SpawnEvent"

type SpawnSystem struct{}

func (s SpawnSystem) Update(w *ecs.World) {
	for _, e := range w.Entities() {
		sp, ok := ecs.TryGet[*component.Spawner](e)
		if !ok {
			continue
		}
		sp.Elapsed++
		if sp.Elapsed >= sp.Interval {
			sp.Elapsed = 0
			w.Emit(ecs.Event{
				Subject: e,
				Verb:    SpawnEvent,
			})
		}
	}
}

func HandleSpawn(w *ecs.World, e ecs.Event) {
	sp, ok := ecs.TryGet[*component.Spawner](e.Subject)
	if !ok {
		return
	}
	w.AddEntity(sp.Factory()...)
}
