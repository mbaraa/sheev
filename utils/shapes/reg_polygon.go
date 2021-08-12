package shapes

import (
	"math"

	"github.com/mbaraa/ligma/models"
)

// RegPolygon represents a regular polygon
type RegPolygon struct {
	numSides       int
	sideLength     float64
	innerSideAngle float64
	vertices       []models.Point2
}

// NewPolygon returns a new Polygon instance
func NewPolygon(numSides int, sideLength float64, vertices ...models.Point2) *RegPolygon {
	return &RegPolygon{
		numSides:       numSides,
		sideLength:     sideLength,
		innerSideAngle: calculateInnerSideAngle(numSides),
		vertices:       vertices,
	}
}

func calculateInnerSideAngle(numSides int) float64 {
	return (math.Pi * 2) / float64(numSides)
}

// GetSideLength return polygon's side length
func (p *RegPolygon) GetSideLength() float64 {
	return p.sideLength
}

// GetNumSides returns number of sides the polygon has
func (p *RegPolygon) GetNumSides() int {
	return p.numSides
}

// GetVertices returns a slice of utils.Point2 that represents the vertices of the polygon
func (p *RegPolygon) GetVertices() []models.Point2 {
	return p.vertices
}

// SetVertices it's written on the box :)
func (p *RegPolygon) SetVertices(v []models.Point2) {
	p.vertices = v
}

// GetInnerSideAngle returns the angle between sides in radians
func (p *RegPolygon) GetInnerSideAngle() float64 {
	return p.innerSideAngle
}
