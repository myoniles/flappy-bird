package ecs

type System interface {
	Update(w World)
}

type Transitioner interface {
	OnEnter(w World)
	OnExit(w World)
}
