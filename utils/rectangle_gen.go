package utils

import (
	"github.com/mbaraa/sheev/utils/shapes"
)

// RectangleGenerator generates a rectangle with the given start and end points
type RectangleGenerator struct {
	bounds  shapes.Bounds
	polygon *shapes.IrRegPolygon
}

// NewRectangleGenerator returns a new RectangleGenerator instance
func NewRectangleGenerator(bounds shapes.Bounds) (r *RectangleGenerator) {
	return &RectangleGenerator{bounds: bounds}
}

// GenerateRectangle generates a rectangle, lol
func (r *RectangleGenerator) GenerateRectangle() *shapes.IrRegPolygon {
	r.calcRectVertices()
	return r.polygon
}

// hmm
func (r *RectangleGenerator) calcRectVertices() {
	r.polygon = shapes.NewIrRegPolygon(4, []shapes.Point2{
		r.bounds.GetMin(),
		{X: r.bounds.GetMax().X, Y: r.bounds.GetMin().Y},
		r.bounds.GetMax(),
		{X: r.bounds.GetMin().X, Y: r.bounds.GetMax().Y},
	}...)
}
