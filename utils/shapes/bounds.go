package shapes

// Point2 represents a point on a 2d plane
type Point2 struct {
	X, Y float64
}

// Bounds represents the top left and bottom right coordinates of a shape
type Bounds struct {
	min *Point2
	max *Point2
}

// NewBounds returns a new Bound instance
func NewBounds(min, max *Point2) *Bounds {
	return &Bounds{min, max}
}

// GetMin returns the top left point of a shape
func (b *Bounds) GetMin() *Point2 {
	return b.min
}

// GetMax returns the bottom right point of a shape
func (b *Bounds) GetMax() *Point2 {
	return b.max
}
