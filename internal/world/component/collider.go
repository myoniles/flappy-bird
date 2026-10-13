package component

import "github.com/myoniles/vulcan/internal/ecs"

const ColliderType ecs.ComponentType = "Collider"

type CollisionLayer uint32

const (
	// 0, care about nothing
	LayerPlayer CollisionLayer = 1 << iota
	LayerObstacle
	LayerBoundary
)

type Collider struct {
	Width, Height int
	Layer         CollisionLayer
	Mask          CollisionLayer
}

func (c *Collider) Type() ecs.ComponentType {
	return ColliderType
}
