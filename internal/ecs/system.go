package ecs

type System interface {
	Update(w *World)
}

type GameState string

type Transitioner interface {
	Transition(w *World)
}

type transition struct {
	To   GameState
	From GameState
}

const AnyState GameState = "ANY"
