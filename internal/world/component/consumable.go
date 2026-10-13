package component

import "github.com/myoniles/vulcan/internal/ecs"

const ConsumableType ecs.ComponentType = "ConsumableType"

type Consumable struct{ Impulse float64 }

func (f Consumable) Type() ecs.ComponentType { return ConsumableType }
