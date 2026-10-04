package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type CollisionSystem struct{}

const CollisionEvent ecs.EventType = "CollisionEvent"

func (c CollisionSystem) Update(entities []*ecs.Entity, q *ecs.EventQueue) {
	for i, e := range entities {
		if !isCollidable(e) {
			continue
		}

		// start at i+1 so each pair is checked once and never against itself
		for _, other := range entities[i+1:] {
			if !isCollidable(other) {
				continue
			}
			// two static things (e.g. a wall and the floor) touching is not news
			if !e.HasComponent("Velocity") && !other.HasComponent("Velocity") {
				continue
			}
			if isColliding(e, other) {
				event := ecs.Event{
					Subject: e,
					Object:  other,
					Verb:    CollisionEvent,
				}
				q.Enqueue(event)
			}
		}
	}
}

func isCollidable(e *ecs.Entity) bool {
	return e.HasComponent("Collider") && e.HasComponent("Transform")
}

func isColliding(e1 *ecs.Entity, e2 *ecs.Entity) bool {
	t1 := e1.GetComponent("Transform").(*component.Transform)
	t2 := e2.GetComponent("Transform").(*component.Transform)

	b1 := e1.GetComponent("Collider").(*component.Collider)
	b2 := e2.GetComponent("Collider").(*component.Collider)

	// Both transforms and colliders are centered, so two boxes overlap on an
	// axis when the distance between centers is less than the sum of their
	// half-extents. Compare doubled values to stay in integer math.
	overlapX := 2*abs(t1.X-t2.X) < b1.Width+b2.Width
	overlapY := 2*abs(t1.Y-t2.Y) < b1.Height+b2.Height

	return overlapX && overlapY
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
