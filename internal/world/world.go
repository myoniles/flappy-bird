package world

import (
	"log"
	"slices"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/input"
	"github.com/myoniles/vulcan/internal/world/system"
)

type World struct {
	entities  []*ecs.Entity
	systems   []ecs.System
	queue     *ecs.EventQueue
	state     ecs.GameState
	nextState ecs.GameState
	input     input.State

	deleted  map[*ecs.Entity]struct{}
	gameLost bool
}

func NewWorld() World {
	w := World{gameLost: false, nextState: ecs.Title}
	queue := ecs.EventQueue{}
	w.deleted = make(map[*ecs.Entity]struct{})

	w.systems = append(w.systems, system.FlapSystem{})
	w.systems = append(w.systems, system.PauseTransition{})
	w.systems = append(w.systems, system.GravitySystem{})
	w.systems = append(w.systems, system.MovementSystem{})
	w.systems = append(w.systems, system.CollisionSystem{})
	w.systems = append(w.systems, system.CleanupSystem{})
	w.systems = append(w.systems, system.SpawnSystem{})
	w.queue = &queue
	return w
}

func (w *World) Update(in input.State) {
	w.input = in
	for _, s := range w.systems {
		s.Update(w)
	}
	w.processEvents()
	w.cleanup()
	w.transition()
}

func (w *World) transition() {
	if w.NextState() == w.State() {
		return
	}
	currTransitioner := system.Transitioners[w.State()]
	for _, t := range currTransitioner {
		t.OnExit(w)
	}

	nextTransitioner := system.Transitioners[w.NextState()]
	for _, t := range nextTransitioner {
		t.OnEnter(w)
	}
	w.state = w.nextState
	log.Printf("Curr %s, Next %s", w.State(), w.NextState())
}

func (w World) Entities() []*ecs.Entity {
	return w.entities
}

func (w World) Query(c ecs.ComponentType) []*ecs.Entity {
	result := make([]*ecs.Entity, 0, len(w.entities))
	for _, v := range w.entities {
		if v.HasComponent(c) {
			result = append(result, v)
		}
	}
	return result
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

	w.entities = slices.DeleteFunc(w.entities, func(e *ecs.Entity) bool {
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

func (w World) State() ecs.GameState {
	return w.state
}

func (w World) NextState() ecs.GameState {
	return w.nextState
}

func (w *World) SetNextState(s ecs.GameState) {
	w.nextState = s
}

func (w *World) processEvents() {
	for {
		event, ok := w.queue.Dequeue()
		if !ok {
			break
		}
		log.Printf("Processing Event: %p. Subject: %p, Object: %s, verb: 5s", &event.Subject, &event.Object, event.Verb)
		if event.Verb == system.GameLossEvent {
			w.gameLost = true
		}

		for _, handler := range system.Handlers[event.Verb] {
			handler(w, event)
		}
	}
}

func (w *World) GetInput() input.State {
	return w.input
}
