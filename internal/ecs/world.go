package ecs

type World interface {
	Entities() []*Entity
	AddEntity(e ...*Entity)
	RemoveEntity(e *Entity)
	Emit(ev Event)
}
