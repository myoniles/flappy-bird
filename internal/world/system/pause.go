package system

import (
	"image/color"
	"log"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/input"
	"github.com/myoniles/vulcan/internal/world/component"
)

func init() {
	Transitioners[ecs.Pause] = append(Transitioners[ecs.Pause], &PauseTransition{}) 
	Handlers[PauseToggleEvent] = append(Handlers[PauseToggleEvent], HandlePauseToggle)
}

const PauseToggleEvent ecs.EventType = "PauseToggle"

type PauseTransition struct {
}

func (t *PauseTransition) OnEnter(w ecs.World) {
	windowConf := config.GetWindowConfig()
	titleText := ecs.NewEntity("Pause")
	// TODO config
	titleText.AddComponent(&component.Text{
		Text:     "Pause",
		FontSize: 67,
		Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
	})
	titleText.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: ecs.Pause})
	w.AddEntity(titleText)

	help := ecs.NewEntity("HelpText")
	// TODO config
	help.AddComponent(&component.Text{
		Text:     "Press Space to start",
		FontSize: 40,
		Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
	})
	help.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight - 100})
	help.AddComponent(&component.StateScoped{Active: ecs.Pause})
	w.AddEntity(help)
}

func (t *PauseTransition) OnExit(w ecs.World) {
}

func (t PauseTransition) Update(w ecs.World) {
	if w.State() != ecs.Pause && w.State() != ecs.PlayState {
		return
	}
	in := w.GetInput()
	log.Print(in)

	if w.GetInput().JustPressed(input.Pause) {
		w.Emit(ecs.Event{
			Verb: PauseToggleEvent,
		})
	}
}

func HandlePauseToggle(w ecs.World, e ecs.Event) {
	if w.State() == ecs.PlayState {
		w.SetNextState(ecs.Pause)
	} else if w.State() == ecs.Pause {
		w.SetNextState(ecs.PlayState)
	}
}
