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
	scoreComp := scorer.GetComponent(component.ScorerType).(*component.Scorer)
	newScore := scoreComp.Score()
	if len(scoreBoards) == 0 {
		log.Panic("There should be at least one scorebaord")
	}
	sb := scoreBoards[0]

	txt := sb.GetComponent(component.TextType).(*component.Text)
	txt.Text = fmt.Sprintf("Score:\n%d", newScore)
}
