package ecs

type EventType string

type Event struct {
	Subject *Entity
	Object  *Entity
	Verb    EventType
}

type EventQueue struct {
	events []Event
}

func (q *EventQueue) Enqueue(e Event) {
	q.events = append(q.events, e)
}

func (q *EventQueue) Dequeue() (Event, bool) {
	if len(q.events) == 0 {
		var zero Event
		return zero, false
	}
	event := q.events[0]
	q.events = q.events[1:]
	return event, true
}

func (q EventQueue) Query(t EventType) *EventQueue {
	newq := EventQueue{}
	for _, e := range q.events {
		if e.Verb == t {
			newq.Enqueue(e)
		}
	}
	return &newq
}

func (e Event) GetEntitiesWithComponent(c ComponentType) []*Entity {
	var comps []*Entity
	if e.Subject.HasComponent(c) {
		comps = append(comps, e.Subject)
	}

	if e.Object.HasComponent(c) {
		comps = append(comps, e.Object)
	}

	return comps
}

type EventHandler func(w *World, e Event)
