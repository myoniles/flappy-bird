package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world"
	"github.com/myoniles/vulcan/internal/world/component"
)

func RenderWorld(w world.World, screen *ebiten.Image) {
	for _, e := range w.Entities() {
		renderSprite(e, screen)
	}
}

func renderSprite(e *ecs.Entity, screen *ebiten.Image) {
	if !e.HasComponent(component.SpriteType) || !e.HasComponent(component.TransformType) {
		return
	}

	s := e.GetComponent(component.SpriteType).(component.Sprite)
	t := e.GetComponent(component.TransformType).(*component.Transform)

	op := &ebiten.DrawImageOptions{}

	spriteImage := ebiten.NewImageFromImage(*s.Img)
	w, h := spriteImage.Bounds().Dx(), spriteImage.Bounds().Dy()
	// Treat (t.X, t.Y) as the entity's center, not its top-left corner.
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	op.Filter = ebiten.FilterLinear

	screen.DrawImage(spriteImage, op)
}
