package ecs

type System interface {
	Update(e []*Entity)
}
