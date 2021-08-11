package models

import (
	"image"

	"github.com/mbaraa/asu_forms/errors"
	"github.com/ungerik/go-cairo"
)

// FormImage represents the source image of a form
type FormImage struct {
	srcImage     image.Image
	workingImage *cairo.Surface
	bounds       *Bounds
}

// NewFormImage returns a new FormImage instance
func NewFormImage(img image.Image) (i *FormImage) {
	return &FormImage{
		srcImage:     img,
		workingImage: cairo.NewSurfaceFromImage(img),
		bounds: NewBounds(
			&Point2{},
			&Point2{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)},
		),
	}
}

// GetBounds returns the image's bounds
func (i *FormImage) GetBounds() *Bounds {
	return i.bounds
}

// ResetChanges resets all of the changes done to the working image surface
func (i *FormImage) ResetChanges() {
	panic(errors.ErrNotImplemented)
}

// GetSurface returns a pointer to the working image surface
// used in field drawing
func (i *FormImage) GetSurface() *cairo.Surface {
	return i.workingImage
}
