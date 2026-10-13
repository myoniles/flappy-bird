package render

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/myoniles/vulcan/internal/ecs"
	"github.com/myoniles/vulcan/internal/world/component"
	"golang.org/x/image/font/gofont/goregular"
)

var textFaceSource *text.GoTextFaceSource = nil

var sprites = newSpriteCache()

func RenderWorld(w ecs.World, screen *ebiten.Image) {
	for _, e := range w.Entities() {
		renderText(e, screen)
		renderSprite(e, screen)
	}
	sprites.sweep()
}

func renderSprite(e *ecs.Entity, screen *ebiten.Image) {
	if !e.HasComponent(component.SpriteType) || !e.HasComponent(component.TransformType) {
		return
	}

	s := e.GetComponent(component.SpriteType).(component.Sprite)
	t := e.GetComponent(component.TransformType).(*component.Transform)

	op := &ebiten.DrawImageOptions{}

	spriteImage := sprites.get(s.Img)
	w, h := spriteImage.Bounds().Dx(), spriteImage.Bounds().Dy()
	// Treat (t.X, t.Y) as the entity's center, not its top-left corner.
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	op.Filter = ebiten.FilterLinear

	screen.DrawImage(spriteImage, op)
}

func renderText(e *ecs.Entity, screen *ebiten.Image) {
	if !e.HasComponent(component.TextType) || !e.HasComponent(component.TransformType) {
		return
	}

	if textFaceSource == nil {
		loadFont()
	}

	textComponent := e.GetComponent(component.TextType).(*component.Text)
	transform := e.GetComponent(component.TransformType).(*component.Transform)

	face := &text.GoTextFace{
		Size:   float64(textComponent.FontSize),
		Source: textFaceSource,
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(transform.X)/2.0, float64(transform.Y)/2.0)
	op.ColorScale.SetWithColor(textComponent.Color)
	op.LineSpacing = float64(textComponent.FontSize)
	text.Draw(screen, textComponent.Text, face, op)

}

func loadFont() {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	textFaceSource = src
}
