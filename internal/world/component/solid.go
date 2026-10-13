package component

import "github.com/myoniles/vulcan/internal/ecs"

const SolidType ecs.ComponentType = "Solid"

// Solid marks an entity that moving entities bounce off of instead of passing
// through.
type Solid struct {
	// Restitution is the fraction of speed kept after a bounce (1 = perfectly elastic).
	Restitution float64
}

func (s *Solid) Type() ecs.ComponentType { return SolidType }
