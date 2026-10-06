package world

import (
	"log"
	"slices"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/input"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/entity"
	"github.com/myoniles/vulcan/internal/world/system"
)

type World struct {
	entities []*ecs.Entity
	systems  []ecs.System
	queue    *ecs.EventQueue

	deleted  map[*ecs.Entity]struct{}
	gameLost bool
}

func NewWorld() World {
	w := World{gameLost: false}
	queue := ecs.EventQueue{}
	w.deleted = make(map[*ecs.Entity]struct{})

	w.queue = &queue
	bird := entity.NewBird()
	w.entities = append(w.entities, bird)
	w.entities = append(w.entities, entity.NewPipePair()...)
	w.entities = append(w.entities, entity.NewCeiling())
	w.entities = append(w.entities, entity.NewFloor())
	w.entities = append(w.entities, entity.NewScoreboard(bird.GetComponent(component.ScorerType).(*component.Scorer)))

	w.systems = append(w.systems, system.GravitySystem{})
	w.systems = append(w.systems, system.MovementSystem{})
	w.systems = append(w.systems, system.CollisionSystem{})
	w.systems = append(w.systems, system.CleanupSystem{})
	w.systems = append(w.systems, system.SpawnSystem{})

	pipeSpawner := ecs.NewEntity("pipe-spawner")
	pipeSpawner.AddComponent(&component.Spawner{Interval: 90, Factory: entity.NewPipePair})
	w.entities = append(w.entities, pipeSpawner)
	return w
}

func (w *World) Update(in input.State) {
	if in.JustPressed(input.Flap) {
		w.queue.Enqueue(ecs.Event{Verb: system.FlapEvent})
	}

	for _, s := range w.systems {
		s.Update(w.entities, w.queue)
	}
	w.processEvents()
	w.cleanup()

	var s string
	for _, e := range w.Entities() {
		s += e.Id()
		s += " "
	}
	log.Print(s)
}

func (w World) Entities() []*ecs.Entity {
	return w.entities
}

func (w *World) AddEntity(e ...*ecs.Entity) {
	w.entities = append(w.entities, e...)
}

// RemoveEntity is a lazy delete so that all events / interactions can be done before removal
func (w *World) RemoveEntity(e *ecs.Entity) {
	w.deleted[e] = struct{}{}
}

func (w *World) cleanup() {
	if len(w.deleted) == 0 {
		return
	}

	w.entities = slices.DeleteFunc( w.entities, func(e *ecs.Entity) bool {
		_, ok := w.deleted[e]
		return ok
	})
	clear(w.deleted)
}

func (w *World) Emit(event ecs.Event) {
	w.queue.Enqueue(event)
}

func (w World) Lost() bool {
	return w.gameLost
}

func (w *World) processEvents() {
	for {
		event, ok := w.queue.Dequeue()
		if !ok {
			break
		}

		if event.Verb == system.GameLossEvent {
			w.gameLost = true
		}

		for _, handler := range system.Handlers[event.Verb] {
			handler(w, event)
		}
	}
}
