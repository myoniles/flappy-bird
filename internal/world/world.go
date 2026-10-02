package world

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/entity"
	"github.com/myoniles/vulcan/internal/world/system"
)

type World struct {
	entities []*ecs.Entity
	systems  []ecs.System
}

func NewWorld() World {
	w := World{}
	w.entities = append(w.entities, entity.NewBird())
	w.entities = append(w.entities, entity.NewPipePair()...)
	w.systems = append(w.systems, system.GravitySystem{})
	w.systems = append(w.systems, system.MovementSystem{})
	return w
}

func (w *World) Update(in any) { // TODO make some sort of input state
	for _, s := range w.systems {
		s.Update(w.entities)
	}
	/*
		w.applyInput(in)
		w.updateAI(in)
		w.integrate(in)
		w.resolveCollisions(in)
		w.processEvents()
		w.cleanUp()
	*/
}

func (w World) Entities() []*ecs.Entity {
	return w.entities
}
