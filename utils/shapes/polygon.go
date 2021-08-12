package shapes

import (
	"github.com/mbaraa/ligma/models"
)

type Polygon interface {
	// GetNumSides returns number of sides the polygon has
	GetNumSides() int
	// GetVertices returns a slice of utils.Point2 that represents the vertices of the polygon
	GetVertices() []models.Point2
	// SetVertices it's written on the box :)
	SetVertices(v []models.Point2)
}
