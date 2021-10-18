package formgen

import (
	"fmt"
	"strings"

	"github.com/mbaraa/sheev/utils/shapes"
)

// MultiLinedTextFieldPlacer represents a text to be placed in a form
type MultiLinedTextFieldPlacer struct {
	numLines        int
	lineHeight      float64
	textFieldPlacer *TextFieldPlacer
}

// NewMultiLinedTextFieldPlacer returns a new MultiLinedTextFieldPlacer instance
// the isRTL optional flag is used to indicate whether a non RTL text is placed in an RTL context
func NewMultiLinedTextFieldPlacer(numLines int, lineHeight float64, textPlacer *TextFieldPlacer) *MultiLinedTextFieldPlacer {
	return &MultiLinedTextFieldPlacer{
		numLines:        numLines,
		lineHeight:      lineHeight,
		textFieldPlacer: textPlacer,
	}
}

// GetBounds returns text field's bounds
func (f *MultiLinedTextFieldPlacer) GetBounds() shapes.Bounds {
	return f.textFieldPlacer.GetBounds()
}

// GetPosition returns text field's position
func (f *MultiLinedTextFieldPlacer) GetPosition() shapes.Point2 {
	return f.textFieldPlacer.GetPosition()
}

// PlaceField draws the text on its parent image, and returns an occurring error
func (f *MultiLinedTextFieldPlacer) PlaceField() error {
	f.partitionText()

	lines := strings.Split(f.textFieldPlacer.text.GetContent(), "\n")
	i := 0.
	for _, line := range lines {
		txt := NewTextFieldPlacer(
			f.textFieldPlacer.text,
			f.textFieldPlacer.fieldXWidth,
			f.textFieldPlacer.fontName,
			f.textFieldPlacer.position,
			f.textFieldPlacer.parent,
			f.textFieldPlacer.isRTL,
		)

		txt.text.SetContent(line)

		txt.position.Y += f.lineHeight * i
		txt.PlaceField()
		i++
	}

	return nil
}

func (f *MultiLinedTextFieldPlacer) partitionText() {
	mutStr := []rune(f.textFieldPlacer.text.GetContent())

	for line := 1; line < f.numLines; line++ {
		fmt.Printf("len: %d\n", len(mutStr))

		field2TextRatio := f.textFieldPlacer.fieldXWidth / f.textFieldPlacer.text.GetXLength()

		if field2TextRatio < 1. {
			endOfLine := int(
				float64(len(mutStr)) *
					(float64(line) *
						(field2TextRatio)))

			for mutStr[endOfLine] != ' ' && endOfLine > 0 { // in order to not split one word over two lines :)
				endOfLine--
			}

			mutStr[endOfLine] = '\n'
		} else {
			return
		}

	}
	f.textFieldPlacer.SetContent(string(mutStr))
}

func (f *MultiLinedTextFieldPlacer) getEndOfLineIndex(stringLength, lineNumber int) int {
	field2TextRatio := f.textFieldPlacer.fieldXWidth / f.textFieldPlacer.text.GetXLength()

	if field2TextRatio < 1. {
		return int(
			float64(stringLength) *
				(float64(lineNumber) *
					(field2TextRatio)))
	}

	return stringLength - 1
}

// GetContent returns the inner text of the current text field
func (f *MultiLinedTextFieldPlacer) GetContent() interface{} {
	return f.textFieldPlacer.GetContent() // string
}

// SetContent sets a new value for the text
func (f *MultiLinedTextFieldPlacer) SetContent(txt interface{}) {
	f.textFieldPlacer.SetContent(txt)
}

// SetPartOfContent sets the value of a character in the string
func (f *MultiLinedTextFieldPlacer) SetPartOfContent(charValue, place interface{}) {
	f.textFieldPlacer.SetPartOfContent(charValue, place)
}
