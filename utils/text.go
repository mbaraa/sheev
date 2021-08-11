package utils

import (
	"fmt"
	"image/color"
	"os"

	"github.com/mbaraa/dsc_logo_generator/logogen"
)

// NewText returns a new logogen.Text instance w/o an error,
// ie when a font error is encountered the Default font is used
func NewText(content string, fgColor color.RGBA64, fontSize float64, fontName string) (txt *logogen.Text) {
	font, err := os.ReadFile(fmt.Sprintf("./res/fonts/%s.ttf", fontName))
	if err != nil {
		font, _ = os.ReadFile("./res/font/Default.ttf")
	}
	txt, _ = logogen.NewText(content, fgColor, fontSize, font)
	return
}
