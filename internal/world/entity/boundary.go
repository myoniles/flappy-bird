package entity

import (
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

const (
	BoundaryPrefix    = "boundary"
	BoundaryThickness = 5
	MaskPrefix        = "rect"
)

// NewVertical returns an invisible wall spanning the full logical height,
// centered on x.
func NewVertical(x int) *ecs.Entity {
	windowConf := config.GetWindowConfig()
	e := ecs.NewEntity(BoundaryPrefix)
	e.AddComponent(&component.Transform{X: x, Y: windowConf.LogicalHeight / 2})
	e.AddComponent(&component.Collider{Width: BoundaryThickness, Height: windowConf.LogicalHeight})
	return e
}

// NewHorizontal returns an invisible wall spanning the full logical width,
// centered on y. Use 0 for the ceiling and LogicalHeight for the floor.
func NewHorizontal(y int) *ecs.Entity {
	windowConf := config.GetWindowConfig()
	e := ecs.NewEntity(BoundaryPrefix)
	e.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: y})
	e.AddComponent(&component.Collider{Width: windowConf.LogicalWidth, Height: BoundaryThickness})
	return e
}

func NewCeiling() *ecs.Entity {

	windowConf := config.GetWindowConfig()

	e := NewHorizontal(0)
	e.AddComponent(&component.Collider{Height: 1, Width: windowConf.LogicalWidth, Layer: component.LayerBoundary})
	return e
}

func NewFloor() *ecs.Entity {
	windowConf := config.GetWindowConfig()
	e := NewHorizontal(windowConf.LogicalHeight)
	e.AddComponent(&component.Collider{Height: 1, Width: windowConf.LogicalWidth, Layer: component.LayerBoundary})
	e.AddComponent(component.Poisoner{})
	return e
}
