package models

// Field represents a field that can be drawn on an image
type Field interface {
	// GetBounds returns field's bounds
	GetBounds() *Bounds
	// GetPosition returns field's position
	GetPosition() *Point2
	// PlaceField draws the field on its parent image, and returns an occurring error
	PlaceField() error
	// CanPlaceField reports whether the field can be placed(w/o overflowing parent) or not
	CanPlaceField() bool
	// SetContent sets a new value for the fields content
	SetContent(value interface{})
}
