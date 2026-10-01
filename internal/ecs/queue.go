package ecs

type Event struct {
	Subject, Object, Verb any
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
