package input

import "github.com/myoniles/vulcan/internal/ecs"

const StateType ecs.ResourceType = "Input"

type Action uint8

const (
	Flap Action = iota
	Pause
)

type State struct {
	just uint32
}

func (s *State) Type() ecs.ResourceType { return StateType }

func (s *State) Trigger(a Action) { s.just |= 1 << a }

func (s State) JustPressed(a Action) bool { return s.just&(1<<a) != 0 }
