package component

import "github.com/myoniles/vulcan/internal/ecs"

const FlapperType ecs.ComponentType = "Flapper"

type Flapper struct{ Impulse float64 }

func (f *Flapper) Type() ecs.ComponentType { return FlapperType }
