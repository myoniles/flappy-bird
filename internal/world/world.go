package world

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/entity"
)

type World struct {
	entities []*ecs.Entity
}

func NewWorld() World {
	w := World{}
	w.entities = append(w.entities, entity.NewBird())
	w.entities = append(w.entities, entity.NewPipePair()...)
	return w
}

func (w *World) Update(in any) { // TODO make some sort of input state
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
