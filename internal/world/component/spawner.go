package component

import "github.com/myoniles/vulcan/internal/ecs"

const SpawnerType ecs.ComponentType = "Spawner"

type Spawner struct {
	Interval int
	Elapsed  int
	Factory  func() []*ecs.Entity
}

func (s *Spawner) Type() ecs.ComponentType { return SpawnerType }
