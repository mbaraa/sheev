package formgen

import (
	"image/color"

	"github.com/mbaraa/ligma/models"
	"github.com/mbaraa/ligma/utils"
	"github.com/mbaraa/ligma/utils/shapes"
)

/*
	DISCLAIMER!!!
	a wise ass human once said:
		"if you cast types like crazy after retrieving them from a data store, then you fucked up storing them properly"
	and apparently I fucked up storing them, maybe I'll fix them in the future, or will I?!

	again I'm sorry for what you're about to see 😅
*/

// CreateFieldPlacer returns a proper field placer according to the given field
func CreateFieldPlacer(field models.Field, parentImage *FormImage) (fp FieldPlacer) {
	switch field.FieldType {
	case models.TextField:
		fp = NewTextFieldPlacer(
			utils.NewText(
				field.Content["text"].(string),
				color.RGBA64{
					R: uint16(field.Content["text_color"].(map[string]interface{})["R"].(float64)),
					G: uint16(field.Content["text_color"].(map[string]interface{})["G"].(float64)),
					B: uint16(field.Content["text_color"].(map[string]interface{})["B"].(float64)),
					A: uint16(field.Content["text_color"].(map[string]interface{})["A"].(float64)),
				},
				field.Content["font_size"].(float64),
				field.Content["font_name"].(string),
			),
			field.Content["x_width"].(float64),
			field.Content["font_name"].(string),
			field.Position,
			parentImage,
			field.Content["is_rtl"].(bool),
		)

	case models.SelectionField:
		fp = NewSelectionFieldPlacer(
			utils.NewPolygonDrawer(
				utils.NewRectangleGenerator(shapes.NewBounds(
					shapes.Point2{
						X: field.Content["shape_vertices"].([]interface{})[0].(map[string]interface{})["X"].(float64),
						Y: field.Content["shape_vertices"].([]interface{})[0].(map[string]interface{})["Y"].(float64),
					},
					shapes.Point2{
						X: field.Content["shape_vertices"].([]interface{})[1].(map[string]interface{})["X"].(float64),
						Y: field.Content["shape_vertices"].([]interface{})[1].(map[string]interface{})["Y"].(float64),
					},
				)).GenerateRectangle(),

				color.RGBA64{
					R: uint16(field.Content["shape_color"].(map[string]interface{})["R"].(float64)),
					G: uint16(field.Content["shape_color"].(map[string]interface{})["G"].(float64)),
					B: uint16(field.Content["shape_color"].(map[string]interface{})["B"].(float64)),
					A: uint16(field.Content["shape_color"].(map[string]interface{})["A"].(float64)),
				},
				utils.NewDrawingOptions(
					field.Content["scale"].(float64),
					field.Content["rotation"].(float64),
					shapes.Point2{},
				),
				parentImage.GetSurface(),
			),
			fixSelectionsTypes(field.Content["selections"].(map[string]interface{})),
			Orientation(field.Content["orientation"].(float64)),
		)
		fp.SetPartOfContent("selection", field.Content["selection"].(string))
	}

	return
}

func fixSelectionsTypes(m map[string]interface{}) (m2 map[string]int) {
	m2 = make(map[string]int)
	for key, val := range m {
		m2[key] = int(val.(float64))
	}
	return
}
