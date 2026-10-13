package world

import (
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/state"
	"github.com/myoniles/vulcan/internal/world/system"
)

func NewFlappyWorld() *ecs.World {
	w := ecs.NewWord(state.Title)

	// Systems
	w.AddSystem(system.FlapSystem{})
	w.AddSystem(system.PauseTransition{})
	w.AddSystem(system.GravitySystem{})
	w.AddSystem(system.MovementSystem{})
	w.AddSystem(system.CollisionSystem{})
	w.AddSystem(system.CleanupSystem{})
	w.AddSystem(system.SpawnSystem{})

	// Transitioners
	w.AddTransitioner(state.Title, &system.TitleTransition{})
	w.AddTransitioner(state.Play, &system.TitleToPlayTransition{})
	w.AddTransitioner(state.End, &system.EndTransition{})
	w.AddTransitioner(state.Pause, &system.PauseTransition{})

	// Event Handlers
	w.OnEvent(system.DestroyEvent, system.HandleOutOfBounds)
	w.OnEvent(system.CollisionEvent, system.HandleCoinScored)
	w.OnEvent(system.GameLossEvent, system.HandleGameLossEvent)
	w.OnEvent(system.FlapEvent, system.HandleFlap)
	w.OnEvent(system.FlapEvent, system.HandleTitleFlap)
	w.OnEvent(system.FlapEvent, system.HandleEndFlap)
	w.OnEvent(system.PauseToggleEvent, system.HandlePauseToggle)
	w.OnEvent(system.ScoreUpEvent, system.HandleUpdateScoreboard)
	w.OnEvent(system.SpawnEvent, system.HandleSpawn)
	return w
}
