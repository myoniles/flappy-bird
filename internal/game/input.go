package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/myoniles/vulcan/internal/input"
)

var keyBindings = map[ebiten.Key]input.Action{
	ebiten.KeySpace:   input.Flap,
	ebiten.KeyArrowUp: input.Flap,
	ebiten.KeyW:       input.Flap,
	ebiten.KeyEscape:  input.Pause,
	ebiten.KeyP:       input.Pause,
}

func poll() input.State {
	var s input.State

	for key, action := range keyBindings {
		if inpututil.IsKeyJustPressed(key) {
			s.Trigger(action)
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		len(inpututil.AppendJustPressedTouchIDs(nil)) > 0 {
		s.Trigger(input.Flap)
	}
	return s

}
