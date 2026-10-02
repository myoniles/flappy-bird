package system

import (
	"github.com/myoniles/vulcan/internal/ecs"
)

type CleanupSystem struct {}

func (c CleanupSystem) Update(entities []*ecs.Entity) {
	for _, e := range entities {
		if !e.HasComponent("Transform") {
			continue
		}
	}
}
