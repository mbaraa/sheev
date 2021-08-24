package models

import (
	"image/color"

	"github.com/mbaraa/ligma/utils/formgen"
	"github.com/mbaraa/ligma/utils/shapes"
)

// FieldType defines the type of the field
type FieldType uint

// FieldType enums
const (
	TextField FieldType = iota
	SelectionField
	ImageField
)

// Field represents a field in a form
type Field struct {
	Name      string        `json:"name"`
	FormName  string        `json:"form_name"`
	FieldType FieldType     `json:"field_type"`
	Position  shapes.Point2 `json:"position"`

	Content map[string]interface{} `json:"content"`
}

// these fuckers are not used any where,
// but they exist to remind me what content each field had 🙂

type textFieldContent struct {
	Text      string       `json:"text"`
	TextColor color.RGBA64 `json:"text_color"`
	FontSize  float64      `json:"font_size"`
	FontName  string       `json:"font_name"`
	IsRTL     bool         `json:"is_rtl"`
}

type selectionFieldContent struct {
	ShapeVertices []shapes.Point2     `json:"shape_vertices"`
	ShapeColor    color.Color         `json:"shape_color"`
	Selections    map[string]int      `json:"selections"`
	Selection     string              `json:"selection"`
	Orientation   formgen.Orientation `json:"orientation"`

	Scale    float64 `json:"scale"`
	Rotation float64 `json:"rotation"`
}
