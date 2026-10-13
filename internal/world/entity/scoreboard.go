package entity

import (
	"fmt"
	"image/color"

	"github.com/myoniles/vulcan/internal/config"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
)

func NewScoreboard() *ecs.Entity {
	windowConf := config.GetWindowConfig()

	e := ecs.NewEntity("scoreboard")
	e.AddComponent(&component.Scoreboard{})
	e.AddComponent(&component.Transform{X: windowConf.LogicalWidth / 2, Y: 100})
	e.AddComponent(&component.Text{
		Text:     fmt.Sprintf("Score:\n%d", 0),
		FontSize: 50,
		Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
	})

	return e
}
