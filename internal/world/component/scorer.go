package component

import "github.com/myoniles/vulcan/internal/ecs"

const ScorerType ecs.ComponentType = "ScorerType"
const ScorerBoardType ecs.ComponentType = "ScoreBoardType"

type Scorer struct {
	score  int
	isLost bool
}

type Scoreboard struct {}

func (s *Scoreboard) Type() ecs.ComponentType {
	return ScorerBoardType
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
