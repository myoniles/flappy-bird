package ecs

import (
	"log"
	"math/rand/v2"
	"strings"
)

type ComponentType string

type ComponentTyper interface {
	Type() ComponentType
}

type Entity struct {
	id         string
	components map[ComponentType]ComponentTyper
}

func NewEntity(id string) *Entity {
	newId := generateID(id)
	return &Entity{
		id: newId,
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

func (e *Entity) Id() string {
	return e.id
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const suffixLength = 5

func generateID(prefix string) string {
	var sb strings.Builder
	sb.Grow(len(prefix) + 1)
	sb.WriteString(prefix)
	sb.WriteRune('-')
	
	sb.Grow(suffixLength)
	for range suffixLength {
		sb.WriteByte(charset[rand.IntN(len(charset))]) 
	}
	
	return sb.String()
}
