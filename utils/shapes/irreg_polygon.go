package shapes

import (
	"github.com/mbaraa/asu_forms/models"
)

type IrRegPolygon struct {
	numSides       int
	sideLength     float64
	innerSideAngle float64
	vertices       []models.Point2
}

func NewIrPolygon(numSides int, sideLength, sideAngle float64) *IrRegPolygon {
	return &IrRegPolygon{
		numSides:       numSides,
		sideLength:     sideLength,
		innerSideAngle: sideAngle,
	}
}

// GetSideLength return polygon's side length
func (p *IrRegPolygon) GetSideLength() float64 {
	return p.sideLength
}

// GetNumSides returns number of sides the polygon has
func (p *IrRegPolygon) GetNumSides() int {
	return p.numSides
}

// GetVertices returns a slice of utils.Point2 that represents the vertices of the polygon
func (p *IrRegPolygon) GetVertices() []models.Point2 {
	return p.vertices
}

// SetVertices it's written on the box :)
func (p *IrRegPolygon) SetVertices(v []models.Point2) {
	p.vertices = v
}
