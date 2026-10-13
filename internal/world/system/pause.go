package system

import (
	"image/color"

	"github.com/myoniles/vulcan/internal/asset"
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

	backgroundMask := ecs.NewEntity("Mask")
	img := asset.NewRect(windowConf.LogicalWidth, windowConf.LogicalHeight, color.RGBA{R: 50, G: 50, B: 50, A: 200})
	backgroundMask.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight / 2})
	backgroundMask.AddComponent(component.NewStaticSprite(&img))
	backgroundMask.AddComponent(&component.StateScoped{Active: ecs.Pause})
	w.AddEntity(backgroundMask)

	titleText := ecs.NewEntity("Pause")
	textConf := config.TitleTextConfig()
	titleText.AddComponent(&component.Text{
		Text:     "Pause",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	titleText.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: ecs.Pause})
	w.AddEntity(titleText)

	help := ecs.NewEntity("HelpText")
	textConf = config.HelperTextConfig()
	help.AddComponent(&component.Text{
		Text:     "Press <esc> to resume",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
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
