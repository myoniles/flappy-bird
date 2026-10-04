package world

import (
	"github.com/myoniles/vulcan/internal/config"
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
}

func NewWorld() World {
	w := World{}
	queue := ecs.EventQueue{}
	w.queue = &queue
	w.entities = append(w.entities, entity.NewBird())
	w.entities = append(w.entities, entity.NewPipePair()...)
	w.systems = append(w.systems, system.GravitySystem{})
	w.systems = append(w.systems, system.MovementSystem{})
	w.systems = append(w.systems, system.CollisionSystem{})
	w.systems = append(w.systems, system.CleanupSystem{})

	windowConf := config.GetWindowConfig()
	w.entities = append(w.entities, entity.NewHorizontal(0))
	w.entities = append(w.entities, entity.NewHorizontal(windowConf.LogicalHeight))
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
}

func (w World) Entities() []*ecs.Entity {
	return w.entities
}

func (w *World) processEvents() {
	for {
		event, ok := w.queue.Dequeue()
		if !ok {
			break
		}

		for _, handler := range system.Handlers[event.Verb] {
			handler(&w.entities, event)
		}
	}
}
