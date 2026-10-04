package component

import "github.com/myoniles/vulcan/internal/ecs"

type Flapper struct{ Impulse float64 }
func (f Flapper) Type() ecs.ComponentType { return "Flapper" }
