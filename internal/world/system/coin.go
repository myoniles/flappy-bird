package system

import (
	"log"

	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func init() {
	Handlers[CollisionEvent] = append(Handlers[CollisionEvent], HandleCoinScored)
}

func HandleCoinScored(w ecs.World, e ecs.Event) {
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
	log.Print("Coin removed")
	log.Printf("%s captured %s. %s's score now at %d", scorer.Id(), coin.Id(), scoreComp.Type(), scoreComp.Score())
}
