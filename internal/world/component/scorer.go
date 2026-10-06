package component

import "github.com/myoniles/vulcan/internal/ecs"

const ScorerType ecs.ComponentType = "ScorerType"

type Scorer struct {
	score  int
	isLost bool
}

func NewScorer() *Scorer {
	return &Scorer{score: 0, isLost: false}
}

func (s *Scorer) Type() ecs.ComponentType {
	return ScorerType
}

func (s *Scorer) Increment() {
	s.score += 1
}

func (s *Scorer) Lost() {
	s.isLost = true
}

func (s *Scorer) Score() int {
	return s.score
}
