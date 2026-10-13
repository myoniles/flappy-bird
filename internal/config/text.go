package config

import "image/color"

type TextConfig struct {
	FontSize int
	Color    color.Color
}

func DefaultTextColor() color.Color {
	return color.RGBA{
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	}
}

func TitleTextConfig() TextConfig{
	return TextConfig {
		FontSize: 67,
		Color: DefaultTextColor(),
	}
}

func HelperTextConfig() TextConfig{
	return TextConfig {
		FontSize: 40,
		Color: DefaultTextColor(),
	}
}
