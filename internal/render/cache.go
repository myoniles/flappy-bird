package render

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// spriteCache maps game-side images to their GPU copies so each image is
// uploaded once instead of every frame.
type spriteCache struct {
	images map[image.Image]*ebiten.Image
	seen   map[image.Image]struct{}
}

func newSpriteCache() *spriteCache {
	return &spriteCache{
		images: make(map[image.Image]*ebiten.Image),
		seen:   make(map[image.Image]struct{}),
	}
}

func (c *spriteCache) get(img image.Image) *ebiten.Image {
	c.seen[img] = struct{}{}
	if eimg, ok := c.images[img]; ok {
		return eimg
	}
	eimg := ebiten.NewImageFromImage(img)
	c.images[img] = eimg
	return eimg
}

// sweep frees every image that was not requested since the last sweep.
func (c *spriteCache) sweep() {
	for img, eimg := range c.images {
		if _, ok := c.seen[img]; !ok {
			eimg.Deallocate()
			delete(c.images, img)
		}
	}
	clear(c.seen)
}
