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
	titleText := ecs.NewEntity("Title")
	textConf := config.TitleTextConfig()
	titleText.AddComponent(&component.Text{
		Text:     "Overcomplicated\nFlappy Bird",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	titleText.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: state.Title})
	w.AddEntity(titleText)

	help := ecs.NewEntity("HelpText")
	textConf = config.HelperTextConfig()
	help.AddComponent(&component.Text{
		Text:     "Press Space to start",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	help.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight - 100})
	help.AddComponent(&component.StateScoped{Active: state.Title})
	w.AddEntity(help)
}
