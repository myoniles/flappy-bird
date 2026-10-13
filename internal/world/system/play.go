package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/entity"
)

type TitleToPlayTransition struct {}

func (t *TitleToPlayTransition) OnEnter(w *ecs.World) {
	bird := entity.NewBird()
	w.AddEntity(bird)
	w.AddEntity(entity.NewPipePair()...)
	w.AddEntity(entity.NewCeiling())
	w.AddEntity(entity.NewFloor())
	w.AddEntity(entity.NewScoreboard())

	pipeSpawner := ecs.NewEntity("pipe-spawner")
	pipeSpawner.AddComponent(&component.Spawner{Interval: 90, Factory: entity.NewPipePair})
	w.AddEntity(pipeSpawner)
}

func (t *TitleToPlayTransition) OnExit(w *ecs.World) {

}
