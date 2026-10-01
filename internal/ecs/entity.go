package ecs

import "log"

type ComponentType string

type ComponentTyper interface {
	Type() ComponentType
}

type Entity struct {
	id         string
	components map[ComponentType]ComponentTyper
}

func NewEntity() *Entity {
	return &Entity{
		components: make(map[ComponentType]ComponentTyper),
	}
}

func (e *Entity) HasComponent(c ComponentType) bool {
	_, ok := e.components[c]
	return ok
}

func (e *Entity) GetComponent(c ComponentType) ComponentTyper {
	if !e.HasComponent(c) {
		log.Fatalf("Component %s trying to access component that it does not have: %s", e.id, c)
	}
	return e.components[c]
}

func (e *Entity) AddComponent(ct ComponentTyper) {
	e.components[ct.Type()] = ct
}
