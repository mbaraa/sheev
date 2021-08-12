package models

import (
	"github.com/mbaraa/ligma/utils/shapes"
)

// Field represents a field that can be drawn on an image
type Field interface {
	// GetBounds returns field's bounds
	GetBounds() *shapes.Bounds
	// GetPosition returns field's position
	GetPosition() *shapes.Point2
	// PlaceField draws the field on its parent image, and returns an occurring error
	PlaceField() error
	// CanPlaceField reports whether the field can be placed(w/o overflowing parent) or not
	CanPlaceField() bool
	// SetContent sets a new value for the fields content
	SetContent(value interface{})
}
