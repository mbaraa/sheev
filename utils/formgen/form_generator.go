package formgen

import (
	"encoding/base64"
	goerr "errors"

	"github.com/mbaraa/sheev/errors"
	"github.com/mbaraa/sheev/models"
)

// FormGenerator generates a form
type FormGenerator struct {
	name           string
	fields         []models.Field
	fieldsPlpacers map[string]FieldPlacer
	formImage      *FormImage
}

// NewFormGenerator returns a new FormGenerator instance
func NewFormGenerator(name string, formImg *FormImage, fields []models.Field) *FormGenerator {
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
func (f *FormGenerator) GetFields() []models.Field {
	return f.fields
}

// MakeForm generates the form with the given fields' data and returns an occurring error
func (f *FormGenerator) MakeForm() (form models.Form, err error) {
	f.createFieldPlacers()

	err = f.placeFields()
	if err != nil {
		return
	}

	// finalize the form image :]
	formImage, status := f.formImage.GetSurface().WriteToPNGStream()
	if status != 0 {
		err = goerr.New(status.String())
		return
	}
	f.formImage.GetSurface().Finish()
	//f.formImage.GetSurface().Destroy()

	form = models.Form{
		Name:       f.name,
		B64FormImg: base64.StdEncoding.EncodeToString(formImage),
	}
	return
}

func (f *FormGenerator) placeFields() error {
	for _, field := range f.fieldsPlpacers {
		if err := field.PlaceField(); err != nil {
			return err
		}
	}
	return nil
}

func (f *FormGenerator) createFieldPlacers() {
	f.fieldsPlpacers = make(map[string]FieldPlacer, len(f.fields))

	for _, field := range f.fields {
		f.fieldsPlpacers[field.Name] = CreateFieldPlacer(field, f.formImage)
	}
}

// AddField adds a field into the form
func (f *FormGenerator) AddField(name string, field FieldPlacer) {
	f.fieldsPlpacers[name] = field
}

// ModifyFieldContent modifies an *existing* field
func (f *FormGenerator) ModifyFieldContent(name string, newContent interface{}) error {
	if _, fieldExists := f.fieldsPlpacers[name]; fieldExists {
		f.fieldsPlpacers[name].SetContent(newContent)
		return nil
	}

	return errors.ErrNoFieldFound
}

// RemoveField removes the first encountered field with the given name
func (f *FormGenerator) RemoveField(name string) {
	delete(f.fieldsPlpacers, name)
}

func (f *FormGenerator) GetFormImage() *FormImage {
	return f.formImage
}
