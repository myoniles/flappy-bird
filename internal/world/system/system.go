package system

import "github.com/myoniles/vulcan/internal/ecs"

type EventHandler func(w ecs.World, e ecs.Event)

var Handlers = map[ecs.EventType][]EventHandler{}

var Transitioners = map[ecs.GameState][]ecs.Transitioner{}
