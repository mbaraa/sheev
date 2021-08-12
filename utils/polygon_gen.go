package utils

import (
	"math"

	"github.com/mbaraa/ligma/utils/shapes"
)

// PolygonGenerator generates the inner regular polygon of the given circumferting circle
type PolygonGenerator struct {
	cirCircle *shapes.Circle
	polygon   *shapes.RegPolygon
	rotation  float64
	toRotate  bool
}

// NewPolygonGenerator returns a new PolygonGenerator instance
func NewPolygonGenerator(numSides int, circumfertingCircle *shapes.Circle, rotationAngle ...float64) (p *PolygonGenerator) {
	p = &PolygonGenerator{
		cirCircle: circumfertingCircle,
		polygon: shapes.NewPolygon(
			numSides,
			calculateSideLength(circumfertingCircle, numSides),
		),
	}
	p.toRotate = rotationAngle != nil
	if p.toRotate {
		p.rotation = rotationAngle[0]
	}

	return p
}

// ResetPolygon resets the polygon
func (p *PolygonGenerator) ResetPolygon() {
	p.SetPolygon(shapes.NewPolygon(3, 1)) // equilateral triangle
}

// SetPolygon lol
func (p *PolygonGenerator) SetPolygon(newPolygon *shapes.RegPolygon) {
	p.polygon = newPolygon
}

// GeneratePolygon generate a polygon with the given stuff
func (p *PolygonGenerator) GeneratePolygon() *shapes.RegPolygon {
	p.calculateVerticesCoordinates()
	return p.polygon
}

// calculateVerticesCoordinates calculates points' polar coordinates (compared to the center of the polygon),
// then converting them into Cartesian coordinates(compared to the center of the polygon),
// and adding the coordinates of the center.
// then voila we got the coordinates of the vertices(heads) 🎉🎉
func (p *PolygonGenerator) calculateVerticesCoordinates() {
	vertices := make([]shapes.Point2, p.polygon.GetNumSides())
	cumAngle := p.polygon.GetInnerSideAngle()
	var x, y, cX, cY float64

	for pi := 0; pi < p.polygon.GetNumSides(); pi++ {
		cX = p.cirCircle.GetPosition().X
		cY = p.cirCircle.GetPosition().Y

		vertices[pi] = shapes.Point2{
			X: cX + getXCartesian(p.cirCircle.GetRadius(), cumAngle),
			Y: cY + getYCartesian(p.cirCircle.GetRadius(), cumAngle),
		}
		if p.toRotate {
			//  x1 = (x0 – xc)cos(θ) – (y0 – yc)sin(θ) + xc
			//  y1 = (x0 – xc)sin(θ) + (y0 – yc)cos(θ) + yc
			x = vertices[pi].X
			y = vertices[pi].Y

			vertices[pi].X = (x-cX)*math.Cos(p.rotation) - (y-cY)*math.Sin(p.rotation) + cX
			vertices[pi].Y = (x-cX)*math.Sin(p.rotation) + (y-cY)*math.Cos(p.rotation) + cY
		}
		cumAngle += p.polygon.GetInnerSideAngle()
	}
	p.polygon.SetVertices(vertices)
}

// from r = S/2sin(pi/n)    |  S is the side length
// we got S = r*2sin(pi/n)  |    for future confusion
func calculateSideLength(c *shapes.Circle, numSides int) float64 {
	return 2 * c.GetRadius() * math.Sin(math.Pi/float64(numSides))
}

func getXCartesian(radius, angle float64) float64 {
	return radius * math.Cos(angle)
}

func getYCartesian(radius, angle float64) float64 {
	return radius * math.Sin(angle)
}
