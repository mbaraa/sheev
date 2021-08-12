package utils

import (
	"image/color"

	"github.com/mbaraa/ligma/utils/shapes"
	"github.com/ungerik/go-cairo"
)

// DrawingOptions represents polygon drawing options
type DrawingOptions struct {
	Scale      float64
	Rotation   float64
	StartPoint shapes.Point2
}

// NewDrawingOptions returns DrawingOptions
func NewDrawingOptions(scale, rotation float64, startingPoint shapes.Point2) *DrawingOptions {
	return &DrawingOptions{
		Scale:      scale,
		Rotation:   rotation,
		StartPoint: startingPoint,
	}
}

// PolygonDrawer draws a polygon on a given cairo.Surface
type PolygonDrawer struct {
	polygon shapes.Polygon
	color   color.Color
	opts    *DrawingOptions
	surface *cairo.Surface
}

// NewPolygonDrawer returns a new PolygonDrawer instance
func NewPolygonDrawer(polygon shapes.Polygon, polygonColor color.Color,
	drawingOptions *DrawingOptions, surface *cairo.Surface) *PolygonDrawer {

	return &PolygonDrawer{
		polygon: polygon,
		color:   polygonColor,
		opts:    drawingOptions,
		surface: surface,
	}
}

func (d *PolygonDrawer) GetPolygon() shapes.Polygon {
	return d.polygon
}

// DrawStroke draws the polygon on the given surface with stroke
func (d *PolygonDrawer) DrawStroke() {
	d.drawPolygon(stroke)
}

// DrawFill draws the polygon on the given surface with fill
func (d *PolygonDrawer) DrawFill() {
	d.drawPolygon(fill)
}

type drawMode uint

const (
	fill = iota
	stroke
)

func (d *PolygonDrawer) drawPolygon(mode drawMode) {
	d.surface.SetSourceRGBA(normalizeRGBA(d.color))
	d.drawPolygonLines()

	switch mode {
	case stroke:
		d.surface.Stroke()
	case fill:
		d.surface.Fill()
	}
}

func (d *PolygonDrawer) drawPolygonLines() {
	for _, vertex := range d.polygon.GetVertices() {
		d.surface.LineTo(
			vertex.X*d.opts.Scale+d.opts.StartPoint.X,
			vertex.Y*d.opts.Scale+d.opts.StartPoint.Y,
		)
	}
	// connect the last vertex to the first one
	d.surface.LineTo(
		d.polygon.GetVertices()[0].X*d.opts.Scale+d.opts.StartPoint.X,
		d.polygon.GetVertices()[0].Y*d.opts.Scale+d.opts.StartPoint.Y,
	)
}

// normalizeRGBA returns a valid cairo rgba color
func normalizeRGBA(rgba color.Color) (float64, float64, float64, float64) {
	r, g, b, a := rgba.RGBA()

	return float64(r) / 255.0,
		float64(g) / 255.0,
		float64(b) / 255.0,
		float64(a) / 255.0
}
