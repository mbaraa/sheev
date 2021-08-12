package shapes

import (
	"math"

	"github.com/mbaraa/ligma/models"
)

// Circle represents a circle, used only to generate a regular polygon,
// so no much to see here :)
type Circle struct {
	radius   float64
	position *models.Point2
}

// NewCircle returns a new Circle instance
func NewCircle(radius, x, y float64) *Circle {
	return &Circle{
		radius:   radius,
		position: &models.Point2{X: x, Y: y},
	}
}

// GetRadius returns the circle's radius
func (c *Circle) GetRadius() float64 {
	return c.radius
}

// GetPosition returns the position of circle's center
func (c *Circle) GetPosition() models.Point2 {
	return *c.position // de-referencing so the position values can't be changed :)
}

// CalcArea returns the circle's area
func (c *Circle) CalcArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

// CalcCircumference returns the circle's circumference
func (c *Circle) CalcCircumference() float64 {
	return 2 * c.radius * math.Pi
}
