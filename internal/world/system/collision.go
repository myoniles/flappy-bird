package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
)

type CollisionSystem struct{}

const CollisionEvent ecs.EventType = "CollisionEvent"

func (c CollisionSystem) Update(entities []*ecs.Entity, q *ecs.EventQueue) {
	for i, e := range entities {
		if !e.HasComponent("Collider") {
			continue
		}

		for j := i; j < len(entities); j++ {
			if !entities[j].HasComponent("Collider") {
				continue
			}
			if isColliding(e, entities[j]) {
				event := ecs.Event{
					Subject: e,
					Object:  entities[j],
					Verb:    CollisionEvent,
				}
				q.Enqueue(event)
			}
		}
	}
}

func isColliding(e1 *ecs.Entity, e2 *ecs.Entity) bool {
	//t1 := e1.GetComponent("Transform").(*component.Transform)
	//t2 := e2.GetComponent("Transform").(*component.Transform)

	//b1 := e1.GetComponent("Collider").(*component.Collider)
	//b2 := e2.GetComponent("Collider").(*component.Collider)

	// Both transforms and colliders are centered.
	// To get the actual boundary locations then,

	return e1 != e2
}
