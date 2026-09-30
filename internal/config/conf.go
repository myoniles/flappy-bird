package config

import (
	"github.com/spf13/viper"
)

type WindowConfig struct {
	Width  int
	Height int
}

func GetWindowConfig() WindowConfig {
	viper.Get("window")
	return WindowConfig{Width: 640, Height: 480}
}
