package component

import "github.com/myoniles/vulcan/internal/ecs"

type Collider struct {
	Width, Height int
}

func (c Collider) Type() ecs.ComponentType {
	return "Collider"
}
