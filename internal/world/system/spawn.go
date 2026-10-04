package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const SpawnEvent ecs.EventType = "SpawnEvent"

func init() {
	Handlers[SpawnEvent] = append(Handlers[SpawnEvent], HandleSpawn)
}

type SpawnSystem struct{}

func (s SpawnSystem) Update(entities []*ecs.Entity, q *ecs.EventQueue) {
	for _, e := range entities {
		if !e.HasComponent(component.SpawnerType) {
			continue
		}
		sp := e.GetComponent(component.SpawnerType).(*component.Spawner)
		sp.Elapsed++
		if sp.Elapsed >= sp.Interval {
			sp.Elapsed = 0
			q.Enqueue(ecs.Event{
				Subject: e,
				Verb:    SpawnEvent,
			})
		}
	}
}

func HandleSpawn(entities *[]*ecs.Entity, e ecs.Event) {
	sp := e.Subject.GetComponent(component.SpawnerType).(*component.Spawner)
	*entities = append(*entities, sp.Factory()...)
}
