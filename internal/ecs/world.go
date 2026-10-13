package ecs

import (
	"log"
	"slices"
)

type GameState string

type World struct {
	entities      []*Entity
	systems       []System
	queue         *EventQueue
	state         GameState
	nextState     GameState
	resources     map[ResourceType]ResourceTyper
	handlers      map[EventType][]EventHandler
	transitioners map[GameState][]Transitioner

	deleted map[*Entity]struct{}
}

func NewWord(initial GameState) *World {
	w := World{nextState: initial}
	// TODO
	w.deleted = make(map[*Entity]struct{})
	w.resources = make(map[ResourceType]ResourceTyper)
	w.handlers = make(map[EventType][]EventHandler)
	w.transitioners = make(map[GameState][]Transitioner)
	w.queue = &EventQueue{}
	return &w
}

func (w *World) AddSystem(s ...System) {
	w.systems = append(w.systems, s...)
}

func (w *World) OnEvent(e EventType, h EventHandler) {
	w.handlers[e] = append(w.handlers[e], h)
}

func (w *World) AddTransitioner(s GameState, t Transitioner) {
	w.transitioners[s] = append(w.transitioners[s], t)
}

func (w *World) Update() {
	for _, s := range w.systems {
		s.Update(w)
	}
	w.processEvents()
	w.transition()
	w.cleanup()
}

func (w *World) transition() {
	if w.NextState() == w.State() {
		return
	}
	currTransitioner := w.transitioners[w.State()]
	for _, t := range currTransitioner {
		t.OnExit(w)
	}

	nextTransitioner := w.transitioners[w.NextState()]
	for _, t := range nextTransitioner {
		t.OnEnter(w)
	}
	w.state = w.nextState
	log.Printf("Curr %s, Next %s", w.State(), w.NextState())
}

func (w World) Entities() []*Entity {
	return w.entities
}

func (w World) Query(c ComponentType) []*Entity {
	result := make([]*Entity, 0, len(w.entities))
	for _, v := range w.entities {
		if v.HasComponent(c) {
			result = append(result, v)
		}
	}
	return result
}

func (w *World) AddEntity(e ...*Entity) {
	w.entities = append(w.entities, e...)
}

// RemoveEntity is a lazy delete so that all events / interactions can be done before removal
func (w *World) RemoveEntity(e *Entity) {
	w.deleted[e] = struct{}{}
}

func (w *World) cleanup() {
	if len(w.deleted) == 0 {
		return
	}

	w.entities = slices.DeleteFunc(w.entities, func(e *Entity) bool {
		_, ok := w.deleted[e]
		return ok
	})
	clear(w.deleted)
}

func (w *World) Emit(event Event) {
	w.queue.Enqueue(event)
}

func (w World) State() GameState {
	return w.state
}

func (w World) NextState() GameState {
	return w.nextState
}

func (w *World) SetNextState(s GameState) {
	w.nextState = s
}

func (w *World) processEvents() {
	for {
		event, ok := w.queue.Dequeue()
		if !ok {
			break
		}
		log.Printf("Processing Event: %p. Subject: %p, Object: %s, verb: 5s", &event.Subject, &event.Object, event.Verb)

		for _, handler := range w.handlers[event.Verb] {
			handler(w, event)
		}
	}
}

// Resource returns the stored ResourceTyper. Like components, resources should
// be stored as pointers so they can be type-asserted and mutated in place.
func (w World) Resource(t ResourceType) ResourceTyper {
	r, ok := w.resources[t]
	if !ok {
		log.Fatalf("World trying to access resource that it does not have: %s", t)
	}
	return r
}

func (w *World) SetResource(r ResourceTyper) {
	w.resources[r.Type()] = r
}
