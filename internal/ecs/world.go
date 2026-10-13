package ecs

type GameState string

type World interface {
	Entities() []*Entity
	Query(c ComponentType) []*Entity
	AddEntity(e ...*Entity)
	RemoveEntity(e *Entity)
	Emit(ev Event)
	State() GameState
	NextState() GameState
	SetNextState(g GameState)
	Resource(t ResourceType) ResourceTyper
	SetResource(r ResourceTyper)
}
