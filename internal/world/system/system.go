package system

import "github.com/myoniles/vulcan/internal/ecs"

type EventHandler func(entities *[]*ecs.Entity, e ecs.Event)

var Handlers = map[ecs.EventType][]EventHandler{}
