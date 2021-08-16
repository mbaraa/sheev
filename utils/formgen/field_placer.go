package formgen

import "github.com/mbaraa/ligma/utils/shapes"

// FieldData represents the field's data, well duh a field has got to have some data :0
type FieldData interface {
	// GetContent returns the content of the current field
	GetContent() interface{}
	// SetContent sets a new value for the fields content
	SetContent(value interface{})
	// SetPartOfContent sets a portion of the content
	//
	// e.g. if we want to modify an entry in a map it would look like this
	// SetPartOfContent(value, key)
	// and for a string
	// SetPartOfContent(charValue, index)
	//
	// believe it or not this little method saved the project
	// from being added to the stack of un-finished projects :]
	SetPartOfContent(content, place interface{})
}

// FieldPlacer places a field on its form
type FieldPlacer interface {
	FieldData
	// PlaceField draws the field on its parent image, and returns an occurring error
	PlaceField() error
}

// fieldPositioning is only used in debugging :}
type fieldPositioning interface {
	// GetBounds returns field's bounds
	GetBounds() *shapes.Bounds
	// GetPosition returns field's position
	GetPosition() *shapes.Point2
}
