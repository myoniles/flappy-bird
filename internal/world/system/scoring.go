package system

import (
	"fmt"
	"log"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func HandleUpdateScoreboard(w *ecs.World, event ecs.Event) {
	scoreBoards := w.Query(component.ScorerBoardType)
	scorer := event.Subject
	scoreComp, ok := ecs.Get[*component.Scorer](scorer)
	if !ok {
		log.Panicf("Entity %s has no Scorer", scorer.Id())
	}
	newScore := scoreComp.Score()
	if len(scoreBoards) == 0 {
		log.Panic("There should be at least one scorebaord")
	}
	sb := scoreBoards[0]

	txt, ok := ecs.Get[*component.Text](sb)
	if !ok {
		log.Panicf("Scoreboard %s has no Text", sb.Id())
	}
	txt.Text = fmt.Sprintf("Score:\n%d", newScore)
}
