package system

import (
	"log"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/state"
)

const ScoreUpEvent ecs.EventType = "CoinScored"

func HandleCoinScored(w *ecs.World, e ecs.Event) {
	if w.State() != state.Play {
		return
	}
	if e.Verb != CollisionEvent {
		return
	}

	if e.Subject == nil || e.Object == nil {
		panic("These should never be nil")
	}

	var scorer *ecs.Entity = nil
	var coin *ecs.Entity = nil
	scorers := e.GetEntitiesWithComponent(component.ScorerType)
	if len(scorers) == 1 {
		scorer = scorers[0]
	}
	coins := e.GetEntitiesWithComponent(component.ConsumableType)
	if len(coins) == 1 {
		coin = coins[0]
	}

	if coin == scorer {
		return
	}

	if coin == nil || scorer == nil {
		log.Printf("Scorer %p\t\tCoin %s", scorer, coin)
		return
	}

	scoreComp := scorer.GetComponent(component.ScorerType).(*component.Scorer)
	scoreComp.Increment()
	w.RemoveEntity(coin)
	log.Printf("%s captured %s. %s's score now at %d", scorer.Id(), coin.Id(), scoreComp.Type(), scoreComp.Score())
	w.Emit(ecs.Event{
		Verb:    ScoreUpEvent,
		Subject: scorer,
	})
}
