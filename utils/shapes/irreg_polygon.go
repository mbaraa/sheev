package shapes

import (
	"github.com/mbaraa/ligma/models"
)

// IrRegPolygon represents an irregular polygon
type IrRegPolygon struct {
	numSides int
	vertices []models.Point2
}

// NewIrRegPolygon returns a new IrRegPolygon instance
func NewIrRegPolygon(numSides int, vertices ...models.Point2) *IrRegPolygon {
	if numSides != len(vertices) {
		return nil
	}

	return &IrRegPolygon{
		numSides: numSides,
		vertices: vertices,
	}
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
