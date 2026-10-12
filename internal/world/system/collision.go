package system

import (
	"log"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func init() {
	//Handlers[CollisionEvent] = append(Handlers[CollisionEvent], HandleBlockBack)
	Handlers[GameLossEvent] = append(Handlers[GameLossEvent], HandleGameLoss)
}

type CollisionSystem struct{}

const CollisionEvent ecs.EventType = "CollisionEvent"
const GameLossEvent ecs.EventType = "GameOver"

func (c CollisionSystem) Update(w ecs.World) {
	if w.State() != ecs.PlayState {
		return
	}
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

func HandleBlockBack(w ecs.World, event ecs.Event) {
	if event.Object.HasComponent(component.VelocityType) {
		v1 := event.Object.GetComponent(component.VelocityType).(*component.Velocity)
		v1.X = 0
		v1.Y = 0

	}

	if event.Subject.HasComponent(component.VelocityType) {
		v1 := event.Subject.GetComponent(component.VelocityType).(*component.Velocity)
		v1.X = 0
		v1.Y = 0
	}
}

func HandleGameLoss(w ecs.World, event ecs.Event) {
	//for _, e := range *entities {
	//	if e.HasComponent(component.VelocityType) {
	//		v := event.Object.GetComponent(component.VelocityType).(*component.Velocity)
	//		v.X = 0
	//		v.Y = 0
	//	}
	//}
	// TODO there is probably a better way to HandleGameLoss then what is in world
	log.Print("GAME OVER")
}
