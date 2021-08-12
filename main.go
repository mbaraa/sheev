package main

import (
	"image/color"

	"github.com/mbaraa/ligma/models"
	"github.com/mbaraa/ligma/utils"
	"github.com/mbaraa/ligma/utils/shapes"
	"github.com/ungerik/go-cairo"
)

func main() {
	c := shapes.NewCircle(2, 1, 1)
	pg := utils.NewPolygonGenerator(6, c)
	p := pg.GeneratePolygon()

	sur := cairo.NewSurface(cairo.FORMAT_ARGB32, 200, 200)
	sur.SetSourceRGBA(1, 1, 1, 1)
	sur.Paint()

	pd := utils.NewPolygonDrawer(p, color.RGBA64{84, 72, 122, 1}, utils.NewDrawingOptions(20, 0, models.Point2{50, 50}), sur)
	pd.DrawStroke()

	sur.WriteToPNG("moo.png")

	sur.Finish()
	sur.Finish()
}
