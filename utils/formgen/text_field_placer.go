package formgen

import (
	"os"

	"github.com/01walid/goarabic"
	"github.com/mbaraa/dsc_logo_generator/logogen"
	"github.com/mbaraa/ligma/utils/shapes"
	"github.com/ungerik/go-cairo"
)

// TextFieldPlacer represents a text to be placed in a form
type TextFieldPlacer struct {
	parent      *FormImage
	bounds      shapes.Bounds
	position    shapes.Point2
	text        logogen.Text
	fontName    string
	fieldXWidth float64
	isRTL       bool
}

// NewTextFieldPlacer returns a new TextFieldPlacer instance
// the isRTL optional flag is used to indicate whether a non RTL text is placed in an RTL context
func NewTextFieldPlacer(text logogen.Text, fieldXWidth float64, fontName string, position shapes.Point2, parent *FormImage, isRTL ...bool) *TextFieldPlacer {
	var isRTL2 bool
	if isRTL != nil {
		isRTL2 = isRTL[0]
	}

	return &TextFieldPlacer{
		text:     text,
		position: position,
		parent:   parent,
		bounds: shapes.NewBounds(
			shapes.Point2{},
			shapes.Point2{X: text.GetXLength(), Y: text.GetFontSize() / 2},
		),
		fieldXWidth: fieldXWidth,
		fontName:    fontName,
		isRTL:       isRTL2,
	}
}

// GetBounds returns text field's bounds
func (f *TextFieldPlacer) GetBounds() shapes.Bounds {
	return f.bounds
}

// GetPosition returns text field's position
func (f *TextFieldPlacer) GetPosition() shapes.Point2 {
	return f.position
}

// PlaceField draws the text on its parent image, and returns an occurring error
func (f *TextFieldPlacer) PlaceField() error {
	f.resizeTextSize()
	f.drawText()

	return nil
}

// resizeText resets the text's size to fit its field
func (f *TextFieldPlacer) resizeTextSize() {
	_, fontSize := f.text.GetXLengthUsingParent(f.fieldXWidth, 1)
	f.text.SetFontSize(fontSize)

}

// drawText draws the
func (f *TextFieldPlacer) drawText() {
	if f.isArabic() {
		f.makeArabic()
	} else if f.isRTL { // I'll fix the mixed text sometime, but for now it is what it is :\
		f.makeRTL()
	}

	f.parent.GetSurface().MoveTo(f.position.X, f.position.Y)
	f.parent.GetSurface().SelectFontFace(f.fontName, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	f.parent.GetSurface().SetSourceRGBA(f.text.GetColorRGBA())
	f.parent.GetSurface().SetFontSize(f.text.GetFontSize())
	f.parent.GetSurface().ShowText(f.text.GetContent())

	f.parent.GetSurface().MoveTo(0, 0) // clean up ish
	f.parent.GetSurface().SetSourceRGBA(0, 0, 0, 1)
	f.parent.GetSurface().SetFontSize(0)
}

// makeArabic sets appropriate settings for the text to be Arabic
func (f *TextFieldPlacer) makeArabic() {
	f.text.SetContent(goarabic.Reverse(goarabic.ToGlyph(f.text.GetContent())))

	f.shiftForRTL()
	f.changeFont("Geeza Pro") // "Gezza Pro" is the most appropriate Arabic font I found :]
}

func (f *TextFieldPlacer) makeRTL() {
	f.shiftForRTL()
	f.changeFont("Default") // "Default.ttf" has all ASCII chars + Arabic
}

func (f *TextFieldPlacer) shiftForRTL() {
	f.position.X -= f.text.GetXLength() // RTL goes brr
}

func (f *TextFieldPlacer) changeFont(fontName string) {
	f.fontName = fontName
	geezaFont, _ := os.ReadFile("./res/fonts/" + fontName + ".ttf")
	f.text.SetFontFamily(geezaFont)
}

// isArabic reports whether the text string is Arabic or not
func (f *TextFieldPlacer) isArabic() bool {
	for _, chr := range f.text.GetContent() {
		if chr >= 0x600 && chr <= 0x6FF {
			return true
		}
	}

	return false
}

// canPlaceField reports whether the text can be placed(w/o overflowing parent) or not
func (f *TextFieldPlacer) canPlaceField() bool {
	return f.bounds.GetMax().X <= f.parent.GetBounds().GetMax().X &&
		f.bounds.GetMax().Y <= f.parent.GetBounds().GetMax().Y
}

// GetContent returns the inner text of the current text field
func (f *TextFieldPlacer) GetContent() interface{} {
	return f.text.GetContent() // string
}

// SetContent sets a new value for the text
func (f *TextFieldPlacer) SetContent(txt interface{}) {
	f.text.SetContent(txt.(string))
	f.bounds = shapes.NewBounds(
		shapes.Point2{},
		shapes.Point2{X: f.text.GetXLength(), Y: f.text.GetFontSize() / 2},
	)
}

// SetPartOfContent sets the value of a character in the string
func (f *TextFieldPlacer) SetPartOfContent(charValue, place interface{}) {
	modifiedStr := []rune(f.text.GetContent())
	modifiedStr[place.(int)%len(modifiedStr)] = charValue.(rune)

	f.SetContent(string(modifiedStr))
}
