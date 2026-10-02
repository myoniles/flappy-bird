package world

import (
	"slices"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/entity"
	"github.com/myoniles/vulcan/internal/world/system"
)

type World struct {
	entities []*ecs.Entity
	systems  []ecs.System
	queue    *ecs.EventQueue
}

func NewWorld() World {
	w := World{}
	queue := ecs.EventQueue{}
	w.queue = &queue
	w.entities = append(w.entities, entity.NewBird())
	w.entities = append(w.entities, entity.NewPipePair()...)
	w.systems = append(w.systems, system.GravitySystem{})
	w.systems = append(w.systems, system.MovementSystem{})
	w.systems = append(w.systems, system.CleanupSystem{})
	return w
}

func (w *World) Update(in any) { // TODO make some sort of input state
	for _, s := range w.systems {
		s.Update(w.entities, w.queue)
	}
	w.processEvent()
	/*
		w.applyInput(in)
		w.updateAI(in) N
		w.integrate(in)
		w.resolveCollisions(in)
		w.processEvents()
		w.cleanUp()
	*/
}

func (w World) Entities() []*ecs.Entity {
	return w.entities
}

func (w *World) processEvent() {
	for {
		event, ok := w.queue.Dequeue()
		if !ok {
			break
		}

		switch event.Verb {
		case system.OutOfBoundsEvent:
			w.entities = slices.DeleteFunc(w.entities, func(e *ecs.Entity) bool {
				return e == event.Subject
			})
		}
	}
}
