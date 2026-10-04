package component

import "github.com/myoniles/vulcan/internal/ecs"

const ColliderType ecs.ComponentType = "Collider"

type Collider struct {
	Width, Height int
}

func (c Collider) Type() ecs.ComponentType {
	return ColliderType
}
