package config

import (
	"github.com/spf13/viper"
)

type WindowConfig struct {
	Width         int
	Height        int
	LogicalWidth  int
	LogicalHeight int
}

func GetWindowConfig() WindowConfig {
	viper.Get("window")
	return WindowConfig{
		Width:         1024,
		Height:        768,
		LogicalWidth:  800,
		LogicalHeight: 600,
	}
}
