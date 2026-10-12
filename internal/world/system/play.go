package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/entity"
)

func init() { Transitioners[ecs.PlayState] = append(Transitioners[ecs.PlayState], &PlayTransition{}) }

type PlayTransition struct {}

func (t *PlayTransition) OnEnter(w ecs.World) {
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

func (t *PlayTransition) OnExit(w ecs.World) {

}
