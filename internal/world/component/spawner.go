package component

import "github.com/myoniles/vulcan/internal/ecs"

type Spawner struct {
	Interval int
	Elapsed  int
	Factory  func() []*ecs.Entity
}

func (s Spawner) Type() ecs.ComponentType { return "Spawner" }
