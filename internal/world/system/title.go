package system

import (
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"github.com/myoniles/vulcan/internal/world/state"
)

type TitleTransition struct{}

func (t *TitleTransition) Transition(w *ecs.World) {
	windowConf := config.GetWindowConfig()
	textConf := config.TitleTextConfig()
	centerX := windowConf.LogicalWidth/2

	// Header
	titleText := ecs.NewEntity("Title")
	titleText.AddComponent(&component.Text{
		Text:     "Overcomplicated\nFlappy Bird",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	titleText.AddComponent(&component.Transform{X: centerX, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: state.Title})
	w.AddEntity(titleText)

	// How to start
	help := ecs.NewEntity("HelpText")
	textConf = config.HelperTextConfig()
	help.AddComponent(&component.Text{
		Text:     "Press Space to start",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	help.AddComponent(&component.Transform{X: centerX, Y: windowConf.LogicalHeight - 100})
	help.AddComponent(&component.StateScoped{Active: state.Title})
	w.AddEntity(help)
}
