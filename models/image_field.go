package models

import (
	"image"

	"github.com/mbaraa/asu_forms/errors"
	"github.com/mbaraa/asu_forms/utils"
)

// ImageField represents an image to be placed in a form
type ImageField struct {
	parent   *FormImage
	bounds   *utils.Bounds
	img      image.Image
	position *utils.Point2
}

// NewImageField returns a new ImageField instance
func NewImageField(name string, img image.Image) *ImageField {
	panic(errors.ErrNotImplemented)
}

// GetBounds returns image field's bounds
func (i *ImageField) GetBounds() *utils.Bounds {
	return i.bounds
}

// GetPosition returns image field's position
func (i *ImageField) GetPosition() *utils.Point2 {
	return i.position
}

// PlaceField draws the image on its parent image, and returns an occurring error
func (i *ImageField) PlaceField() error {
	return errors.ErrNotImplemented
}

// CanPlaceField reports whether the image can be placed(w/o overflowing parent) or not
func (i *ImageField) CanPlaceField() bool {
	panic(errors.ErrNotImplemented)
}

// SetContent sets a new value for the image
func (i *ImageField) SetContent(value interface{}) {
	panic(errors.ErrNotImplemented)
}
