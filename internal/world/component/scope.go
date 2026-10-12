package component

import "github.com/myoniles/vulcan/internal/ecs"

const StateScopedType ecs.ComponentType = "StateScoped"

type StateScoped struct {
	Active ecs.GameState
}

func (ss *StateScoped) Type() ecs.ComponentType {
	return StateScopedType
}
