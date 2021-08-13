package models

import (
	"github.com/mbaraa/ligma/utils"
	"github.com/mbaraa/ligma/utils/shapes"
)

// SelectionField represents a selection(highlighter) field ho ho ho to be placed in a form
type SelectionField struct {
	drawer           *utils.PolygonDrawer
	bounds           *shapes.Bounds
	originalVertices []shapes.Point2

	selections  map[string]int
	selection   string
	orientation Orientation
}

// Orientation represents selection field's orientation
type Orientation int

// Orientations :)
const (
	VerticalSelection Orientation = iota
	HorizontalSelection
)

// NewSelectionField returns a new SelectionField instance
func NewSelectionField(polygonDrawer *utils.PolygonDrawer, selections map[string]int, orientation Orientation) *SelectionField {

	return (&SelectionField{
		drawer:           polygonDrawer,
		orientation:      orientation,
		selections:       selections,
		originalVertices: polygonDrawer.GetPolygon().GetVertices(),
	}).initBounds()
}

// getMaxPolygonPoint returns the maximum x and y coordinates of the given polygon
// it's weird, but it works :]
func getMaxPolygonPoint(polygon shapes.Polygon) *shapes.Point2 {
	xMax, yMax := polygon.GetVertices()[0].X, polygon.GetVertices()[0].Y
	for _, v := range polygon.GetVertices() {
		if v.X > xMax {
			xMax = v.X
		}
		if v.Y > yMax {
			yMax = v.Y
		}
	}

	return &shapes.Point2{X: xMax, Y: yMax}
}

func (s *SelectionField) initBounds() *SelectionField {
	s.bounds = shapes.NewBounds(
		&s.drawer.GetPolygon().GetVertices()[0],
		getMaxPolygonPoint(s.drawer.GetPolygon()),
	)

	return s
}

// GetBounds returns image field's bounds
func (s *SelectionField) GetBounds() *shapes.Bounds {
	return s.bounds
}

// GetPosition returns image field's position
func (s *SelectionField) GetPosition() *shapes.Point2 {
	return nil
}

// PlaceField draws the selection box on its parent image, and returns an occurring error
func (s *SelectionField) PlaceField() error {
	s.fixPositions()
	s.drawer.DrawFill()

	return nil
}

func (s *SelectionField) fixPositions() {
	newVertices := make([]shapes.Point2, s.drawer.GetPolygon().GetNumSides())
	s.drawer.GetPolygon().SetVertices(s.originalVertices)

	for index, vertex := range s.drawer.GetPolygon().GetVertices() {
		newVertices[index] = vertex

		switch s.orientation {
		case VerticalSelection:
			if s.selections[s.selection] > 1 {
				newVertices[index].Y += float64(s.selections[s.selection]-1) * (s.bounds.GetMax().Y - s.bounds.GetMin().Y)
			}
		case HorizontalSelection:
			if s.selections[s.selection] > 1 {
				newVertices[index].X += float64(s.selections[s.selection]-1) * (s.bounds.GetMax().X - s.bounds.GetMin().X)
			}
		}
	}

	s.drawer.GetPolygon().SetVertices(newVertices)
}

// GetContent returns the content of the current field
func (s *SelectionField) GetContent() interface{} {
	return s.selections // map[string]int
}

// SetContent sets the position of the selection box depending on the box name
func (s *SelectionField) SetContent(selection interface{}) {
	if _, selectionExisits := s.selections[selection.(string)]; selectionExisits {
		s.selection = selection.(string)
	}
}
