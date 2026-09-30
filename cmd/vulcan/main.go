package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	cfg "github.com/myoniles/vulcan/internal/config"
	game "github.com/myoniles/vulcan/internal/game"
)

func main() {
	windowConf := cfg.GetWindowConfig()
	ebiten.SetWindowSize(windowConf.Width , windowConf.Height)
	ebiten.SetWindowTitle("Vulcan")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
