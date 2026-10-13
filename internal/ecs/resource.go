package ecs

type ResourceType string

// ResourceTyper is a world-wide singleton (e.g. input, time) that is not tied
// to any entity.
type ResourceTyper interface {
	Type() ResourceType
}
