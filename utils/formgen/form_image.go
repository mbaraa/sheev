package formgen

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/mbaraa/sheev/errors"
	"github.com/mbaraa/sheev/utils/shapes"
	"github.com/ungerik/go-cairo"
)

// FormImage represents the source image of a form
type FormImage struct {
	srcImage     image.Image
	workingImage *cairo.Surface
	bounds       shapes.Bounds
}

// NewFormImage returns a new FormImage instance
func NewFormImage(img image.Image) (i *FormImage) {
	return &FormImage{
		srcImage:     img,
		workingImage: cairo.NewSurfaceFromImage(img),
		bounds: shapes.NewBounds(
			shapes.Point2{},
			shapes.Point2{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)},
		),
	}
}

// NewFormImageFromB64Image returns a new FormImage instance using a base64 image string
func NewFormImageFromB64Image(b64 string) *FormImage {
	img0, _ := base64.StdEncoding.DecodeString(b64)
	bb := bytes.NewReader(img0)
	img1, _ := png.Decode(bb)

	return NewFormImage(img1)
}

// GetBounds returns the image's bounds
func (i *FormImage) GetBounds() shapes.Bounds {
	return i.bounds
}

// ResetChanges resets all of the changes done to the working image surface
func (i *FormImage) ResetChanges() {
	panic(errors.ErrNotImplemented)
}

// GetSurface returns a pointer to the working image surface
// used in field drawing
func (i *FormImage) GetSurface() *cairo.Surface {
	return i.workingImage
}

// GetImage returns the image in base64
func (i *FormImage) GetImage() string {
	img := bytes.NewBuffer([]byte(""))
	err := png.Encode(img, i.srcImage)

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(img.Bytes())
}
