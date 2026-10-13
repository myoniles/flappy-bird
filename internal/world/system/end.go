package system

import (
	"image/color"

	"github.com/myoniles/vulcan/internal/asset"
	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func init() {
	Transitioners[ecs.End] = append(Transitioners[ecs.End], &EndTransition{})
	Handlers[GameLossEvent] = append(Handlers[GameLossEvent], HandleGameLossEvent)
}

type EndTransition struct {
}

func (t *EndTransition) OnEnter(w ecs.World) {
	windowConf := config.GetWindowConfig()

	backgroundMask := ecs.NewEntity("Mask")
	img := asset.NewRect(windowConf.LogicalWidth, windowConf.LogicalHeight, color.RGBA{R: 50, G: 50, B: 50, A: 200})
	backgroundMask.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight / 2})
	backgroundMask.AddComponent(component.NewStaticSprite(&img))
	backgroundMask.AddComponent(&component.StateScoped{Active: ecs.End})
	w.AddEntity(backgroundMask)

	titleText := ecs.NewEntity("End")
	textConf := config.TitleTextConfig()
	titleText.AddComponent(&component.Text{
		Text:     "BOO YOU STINK",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	titleText.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 200})
	titleText.AddComponent(&component.StateScoped{Active: ecs.End})
	w.AddEntity(titleText)


	help := ecs.NewEntity("HelpText")
	textConf = config.HelperTextConfig()
	help.AddComponent(&component.Text{
		Text:     "Press Space to restart",
		FontSize: textConf.FontSize,
		Color:    textConf.Color,
	})
	help.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: windowConf.LogicalHeight - 100})
	help.AddComponent(&component.StateScoped{Active: ecs.End})
	w.AddEntity(help)
}

func (t *EndTransition) OnExit(w ecs.World) {
	for _, e := range w.Entities() {
		w.RemoveEntity(e)
	}
}

func HandleGameLossEvent(w ecs.World, e ecs.Event) {
	w.SetNextState(ecs.End)
}
