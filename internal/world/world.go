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
	w.AddSystem(system.CleanupSystem{})
	w.AddSystemIn(state.Play, system.GravitySystem{})
	w.AddSystemIn(state.Play, system.MovementSystem{})
	w.AddSystemIn(state.Play, system.CollisionSystem{})
	w.AddSystemIn(state.Play, system.SpawnSystem{})
	w.AddSystemIn(state.Play, system.PauseSystem{})
	w.AddSystemIn(state.Pause, system.PauseSystem{})

	// Transitioners
	w.AddTransitioner(ecs.AnyState, state.Title, &system.TitleTransition{})
	w.AddTransitioner(state.Title, state.Play, &system.TitleToPlayTransition{})
	w.AddTransitioner(state.Play, state.End, &system.EndTransition{})
	w.AddTransitioner(state.End, state.Play, &system.ResetGameTransition{})
	w.AddTransitioner(state.Play, state.Pause, &system.PauseTransition{})
	w.AddTransitioner(state.End, state.Play, &system.TitleToPlayTransition{})

	// Event Handlers
	w.OnEvent(system.DestroyEvent, system.HandleOutOfBounds)
	w.OnEvent(system.CollisionEvent, system.HandleCoinScored)
	w.OnEvent(system.CollisionEvent, system.HandleBounce)
	w.OnEvent(system.CollisionEvent, system.HandlePoison)
	w.OnEvent(system.GameLossEvent, system.HandleGameLossEvent)
	w.OnEvent(system.ScoreUpEvent, system.HandleUpdateScoreboard)
	w.OnEvent(system.SpawnEvent, system.HandleSpawn)

	// Input
	w.OnEventIn(state.Title, system.FlapEvent, system.HandleStart)
	w.OnEventIn(state.End, system.FlapEvent, system.HandleStart)
	w.OnEventIn(state.Play, system.FlapEvent, system.HandleFlap)
	w.OnEventIn(state.Play, system.PauseToggleEvent, system.HandlePause)
	w.OnEventIn(state.Pause, system.PauseToggleEvent, system.HandleResume)
	return w
}
