package world

import "github.com/myoniles/vulcan/internal/world/entity"

type World struct {
	entities []*Entity
}

func NewWord() World {
	w := World{}
	w.entities = append(w.entities, entity.NewBird())

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
