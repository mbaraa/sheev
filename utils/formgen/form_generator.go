package formgen

import (
	"fmt"

	"github.com/mbaraa/ligma/errors"
)

// FormGenerator generates a form
type FormGenerator struct {
	name      string
	fields    map[string]FieldPlacer
	formImage *FormImage
}

// NewFormGenerator returns a new FormGenerator instance
func NewFormGenerator(name string, formImg *FormImage, fields map[string]FieldPlacer) *FormGenerator {
	return &FormGenerator{
		name:      name,
		formImage: formImg,
		fields:    fields,
	}
}

// GetName returns form's name
func (f *FormGenerator) GetName() string {
	return f.name
}

// GetFields returns fields, lol
func (f *FormGenerator) GetFields() map[string]FieldPlacer {
	return f.fields
}

// MakeForm generates the form with the given fields' data and returns an occurring error
func (f *FormGenerator) MakeForm() ([]byte, error) {
	err := f.placeFields()
	if err != nil {
		return nil, err
	}

	// finalize the form image :]
	form, _ := f.formImage.GetSurface().WriteToPNGStream()
	f.formImage.GetSurface().Finish()
	//f.formImage.GetSurface().Destroy()

	return form, nil
}

func (f *FormGenerator) placeFields() error {
	for _, field := range f.fields {
		if err := field.PlaceField(); err != nil {
			return err
		}
	}
	fmt.Println("foook", f.fields["اسم الطالب"].GetContent())
	return nil
}

// AddField adds a field into the form
func (f *FormGenerator) AddField(name string, field FieldPlacer) {
	f.fields[name] = field
}

// ModifyFieldContent modifies an *existing* field
func (f *FormGenerator) ModifyFieldContent(name string, newContent interface{}) error {
	if _, fieldExists := f.fields[name]; fieldExists {
		f.fields[name].SetContent(newContent)
		return nil
	}

	return errors.ErrNoFieldFound
}

// RemoveField removes the first encountered field with the given name
func (f *FormGenerator) RemoveField(name string) {
	delete(f.fields, name)
}

func (f *FormGenerator) GetFormImage() *FormImage {
	return f.formImage
}
