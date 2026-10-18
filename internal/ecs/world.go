package ecs

import (
	"log"
	"slices"
)

type World struct {
	entities      []*Entity
	systems       map[GameState][]System
	queue         *EventQueue
	state         GameState
	nextState     GameState
	resources     map[ResourceType]ResourceTyper
	handlers      map[handlerKey][]EventHandler
	transitioners map[transition][]Transitioner

	deleted map[*Entity]struct{}
}

func NewWorld(initial GameState) *World {
	w := World{nextState: initial}
	w.systems = make(map[GameState][]System)
	w.deleted = make(map[*Entity]struct{})
	w.resources = make(map[ResourceType]ResourceTyper)
	w.handlers = make(map[handlerKey][]EventHandler)
	w.transitioners = make(map[transition][]Transitioner)
	w.queue = &EventQueue{}
	return &w
}

// AddSystem is the Same as AddSystemIn(ANYSTATE, sys)
func (w *World) AddSystem(s ...System) {
	w.AddSystemIn(AnyState, s...)
}

func (w *World) AddSystemIn(g GameState, s ...System) {
	w.systems[g] = append(w.systems[g], s...)

}

// OnEvent is the same as OnEventIn(AnyState, e, h)
func (w *World) OnEvent(e EventType, h EventHandler) {
	w.OnEventIn(AnyState, e, h)
}

// OnEventIn registers h to run for events of type e, but only while the world
// is in state g.
func (w *World) OnEventIn(g GameState, e EventType, h EventHandler) {
	k := handlerKey{State: g, Event: e}
	w.handlers[k] = append(w.handlers[k], h)
}

func (w *World) AddTransitioner(from, to GameState, t Transitioner) {
	tr := transition{From: from, To: to}
	w.transitioners[tr] = append(w.transitioners[tr], t)
}

func (w *World) Update() {
	w.updateSystems()
	w.processEvents()
	w.transition()
	w.cleanup()
}

func (w *World) updateSystems() {
	sys := w.systems[AnyState]
	sys = append(sys, w.systems[w.state]...)
	for _, s := range sys {
		s.Update(w)
	}
}

func (w *World) transition() {
	if w.NextState() == w.State() {
		return
	}
	tr := transition{
		From: w.State(),
		To:   w.NextState(),
	}
	currTransitioner := w.transitioners[tr]
	anyToTransitioners := w.transitioners[transition{From: AnyState, To: w.NextState()}]

	currTransitioner = append(currTransitioner, anyToTransitioners...)
	for _, t := range currTransitioner {
		t.Transition(w)
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
		log.Printf("Processing Event: %p. Subject: %p, Object: %s, verb: %s", &event.Subject, &event.Object, event.Verb)

		// Handlers match on the current state, not the next one, so a handler
		// that requests a transition can't trigger another state's handlers
		// for the same event.
		for _, handler := range w.handlers[handlerKey{State: AnyState, Event: event.Verb}] {
			handler(w, event)
		}
		for _, handler := range w.handlers[handlerKey{State: w.state, Event: event.Verb}] {
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
