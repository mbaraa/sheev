package models

import (
	"github.com/mbaraa/ligma/errors"
)

// Form represents a form to be filled :)
type Form struct {
	name      string
	fields    map[string]Field
	formImage *FormImage
}

// NewForm returns a new Form instance
func NewForm(name string, formImg *FormImage, fields map[string]Field) *Form {
	return &Form{
		name:      name,
		formImage: formImg,
		fields:    fields,
	}
}

// GetName returns form's name
func (f *Form) GetName() string {
	return f.name
}

// MakeForm generates the form with the given fields' data and returns an occurring error
func (f *Form) MakeForm() ([]byte, error) {
	for _, form := range f.fields {
		if err := form.PlaceField(); err != nil {
			return nil, err
		}
	}

	// finalize the form image :]
	form, _ := f.formImage.GetSurface().WriteToPNGStream()
	f.formImage.GetSurface().Finish()
	// f.formImage.GetSurface().Destroy()

	return form, nil
}

// AddField adds a field into the form
func (f *Form) AddField(name string, field Field) {
	f.fields[name] = field
}

// ModifyFieldContent modifies an *existing* field
func (f *Form) ModifyFieldContent(name string, newContent interface{}) error {
	if _, fieldExists := f.fields[name]; fieldExists {
		f.fields[name].SetContent(newContent)
		return nil
	}

	return errors.ErrNoFieldFound
}

// RemoveField removes the first encountered field with the given name
func (f *Form) RemoveField(name string) {
	panic(errors.ErrNotImplemented)
}
