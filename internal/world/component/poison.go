package component

import "github.com/myoniles/vulcan/internal/ecs"

const PoisonerType ecs.ComponentType = "Poisoner"

type Poisoner struct{}

func (p Poisoner) Type() ecs.ComponentType {
	return PoisonerType
}

const PoisoneeType ecs.ComponentType = "Poisonee"

type Poisonee struct{}

func (p Poisonee) Type() ecs.ComponentType {
	return PoisoneeType
}
