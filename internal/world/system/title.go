package system

import (
	"image/color"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func init() { Transitioners[ecs.Title] = append(Transitioners[ecs.Title], &TitleTransition{}) }

type TitleTransition struct {
}

func (t *TitleTransition) OnEnter(w ecs.World) {
	windowConf := config.GetWindowConfig()
	titleText := ecs.NewEntity("Title")
	// TODO config
	titleText.AddComponent(&component.Text{
		Text:     "Overcomplicated\nFlappy Bird",
		FontSize: 67,
		Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
	})
	titleText.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: ecs.Title})
	w.AddEntity(titleText)

	help := ecs.NewEntity("HelpText")
	// TODO config
	help.AddComponent(&component.Text{
		Text:     "Press Space to start",
		FontSize: 40,
		Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
	})
	help.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight - 100})
	help.AddComponent(&component.StateScoped{Active: ecs.Title})
	w.AddEntity(help)
}

func (t *TitleTransition) OnExit(w ecs.World) {
}
