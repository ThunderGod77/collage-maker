package ImageManipulation

import (
	"image"

	"github.com/disintegration/imaging"
)

func ResizeImage(a, b, c image.Image) (image.Image, image.Image, image.Image) {
	aWidth := a.Bounds().Size().X
	aHeight := a.Bounds().Size().Y
	bWidth := b.Bounds().Size().X
	bHeight := b.Bounds().Size().Y
	cWidth := c.Bounds().Size().X
	cHeight := c.Bounds().Size().Y

	avgWidth := (aWidth + bWidth + cWidth) / 3
	avgHeight := (aHeight + bHeight + cHeight) / 3

	a = imaging.Resize(a, avgWidth, avgHeight, imaging.Lanczos)
	b = imaging.Resize(b, avgWidth, avgHeight, imaging.Lanczos)
	c = imaging.Resize(c, avgWidth, avgHeight, imaging.Lanczos)

	return a, b, c
}
