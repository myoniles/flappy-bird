package system

import (
	"math"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

type CollisionSystem struct{}

const CollisionEvent ecs.EventType = "CollisionEvent"
const GameLossEvent ecs.EventType = "GameOver"

func (c CollisionSystem) Update(w *ecs.World) {
	entities := w.Entities()
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
			if !e.HasComponent(component.VelocityType) && !other.HasComponent(component.VelocityType) {
				continue
			}
			if isColliding(e, other) {
				event := ecs.Event{
					Subject: e,
					Object:  other,
					Verb:    CollisionEvent,
				}
				w.Emit(event)

				if hasPoisoned(e, other) {
					w.Emit(ecs.Event{Verb: GameLossEvent})
				}
			}
		}
	}
}

func isCollidable(e *ecs.Entity) bool {
	return e.HasComponent(component.ColliderType) && e.HasComponent(component.TransformType)
}

func isColliding(e1 *ecs.Entity, e2 *ecs.Entity) bool {
	b1 := e1.GetComponent(component.ColliderType).(*component.Collider)
	b2 := e2.GetComponent(component.ColliderType).(*component.Collider)

	// Check if they *can* even collide based on their layers
	if b1.Mask&b2.Layer == 0 && b2.Mask&b1.Layer == 0 {
		return false
	}

	t1 := e1.GetComponent(component.TransformType).(*component.Transform)
	t2 := e2.GetComponent(component.TransformType).(*component.Transform)

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

func hasPoisoned(e, other *ecs.Entity) bool {
	if e.HasComponent(component.PoisoneeType) && other.HasComponent(component.PoisonerType) {
		return true
	}
	if e.HasComponent(component.PoisonerType) && other.HasComponent(component.PoisoneeType) {
		return true
	}
	return false
}

// HandleBounce pushes a moving entity out of any Solid it overlaps and reflects
// its velocity along the axis it penetrated least. Collisions that don't
// involve a Solid are ignored.
func HandleBounce(w *ecs.World, event ecs.Event) {
	mover, solid := event.Subject, event.Object
	if !solid.HasComponent(component.SolidType) {
		mover, solid = solid, mover
	}
	if !solid.HasComponent(component.SolidType) || !mover.HasComponent(component.VelocityType) {
		return
	}

	s := solid.GetComponent(component.SolidType).(*component.Solid)
	v := mover.GetComponent(component.VelocityType).(*component.Velocity)
	mt := mover.GetComponent(component.TransformType).(*component.Transform)
	st := solid.GetComponent(component.TransformType).(*component.Transform)
	mc := mover.GetComponent(component.ColliderType).(*component.Collider)
	sc := solid.GetComponent(component.ColliderType).(*component.Collider)

	// Doubled overlap on each axis, matching isColliding's integer math.
	overlapX := mc.Width + sc.Width - 2*abs(mt.X-st.X)
	overlapY := mc.Height + sc.Height - 2*abs(mt.Y-st.Y)

	// Round the push up so the boxes no longer overlap next frame; otherwise
	// the velocity would flip back and forth while the mover stays stuck.
	if overlapY <= overlapX {
		dir := sign(mt.Y - st.Y)
		mt.Y += dir * (overlapY + 1) / 2
		v.Y = float64(dir) * math.Abs(v.Y) * s.Restitution
	} else {
		dir := sign(mt.X - st.X)
		mt.X += dir * (overlapX + 1) / 2
		v.X = float64(dir) * math.Abs(v.X) * s.Restitution
	}
}

// sign returns -1 for negative x and 1 otherwise, so dead-center overlaps
// still get pushed somewhere.
func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
