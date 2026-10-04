package entity

import (
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func NewVertical(x int) *ecs.Entity {
	e := ecs.NewEntity("spawner")
	windowConf := config.GetWindowConfig()
	e.AddComponent(component.Transform{X: x, Y:0})
	e.AddComponent(component.Collider{Width: 5, Height: windowConf.LogicalHeight})
	return e
}
