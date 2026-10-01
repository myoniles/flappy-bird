package entity

import "github.com/myoniles/vulcan/internal/world/component"

type Pipe struct {
	Top    halfPipe
	Bottom halfPipe
}

type halfPipe struct {
	component.Sprite
	component.Position
}
