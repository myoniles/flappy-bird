package input

type Action uint8

const (
	Flap Action = iota
	Pause
)

type State struct {
	just uint32
}

func (s *State) Trigger(a Action) { s.just |= 1 << a }

func (s State) JustPressed(a Action) bool { return s.just&(1<<a) != 0 }
