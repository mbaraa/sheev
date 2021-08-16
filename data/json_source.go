package data

import (
	"encoding/json"
	"os"
	"path"

	"github.com/mbaraa/ligma/errors"
	"github.com/mbaraa/ligma/models"
)

// JSONSource represent a JSON source for the stuff
type JSONSource struct {
	forms0  []*models.Form
	fields0 []*models.Field

	// maps, much speed wow!
	forms map[string]*models.Form // formName -> form

	// I have no idea why I did this one :\
	fields map[string]map[string]*models.Field // parentFormName -> fieldName -> field
}

// NewJSONSource returns a new JSONSource instance
// its complexity is O(fuck) but it's better than O(n) on each field/form fetch :]
func NewJSONSource(jsonDir string) (j *JSONSource) {
	j = new(JSONSource).
		loadJSONFiles(jsonDir)

	// malloc
	j.forms = make(map[string]*models.Form)
	j.fields = make(map[string]map[string]*models.Field)
	var formFields []*models.Field

	for _, form := range j.forms0 {
		j.forms[form.Name] = form
		// more malloc
		formFields = make([]*models.Field, 0)
		j.fields[form.Name] = make(map[string]*models.Field)

		for _, field := range j.fields0 {
			if field.FormName == form.Name { // add form's fields to its fields slice :]
				formFields = append(formFields, field)
			}
			// I have no idea why I did this :\
			j.fields[form.Name][field.Name] = field
		}
		// bla bla bla
		j.forms[form.Name].Fields = formFields
	}

	return j
}

func (j *JSONSource) loadJSONFiles(jsonDir string) *JSONSource {
	forms0, err := os.ReadFile(path.Join(jsonDir, "forms.json"))
	if err != nil {
		panic("ahem, hello where is the `forms.json` file?\nerr: " + err.Error())
	}

	fields0, err := os.ReadFile(path.Join(jsonDir, "fields.json"))
	if err != nil {
		panic("ahem, hello where is the `fields.json` file?\nerr: " + err.Error())
	}

	err = json.Unmarshal(forms0, &j.forms0)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fields0, &j.fields0)

	if err != nil {
		panic(err)
	}

	return j
}

// ExistsByName reports whether the FormGenerator exists or not, and an occurring error
func (j *JSONSource) ExistsByName(name string) bool {
	_, formExists := j.forms[name]
	return formExists
}

// Get returns a form depending on its name, and an occurring error
func (j *JSONSource) Get(name string) (*models.Form, error) {
	if form, formExists := j.forms[name]; formExists {
		return form, nil
	}
	return nil, errors.ErrNoFormFound
}

// GetAll returns all available forms, and an occurring error
func (j *JSONSource) GetAll() ([]*models.Form, error) {
	return j.forms0, nil
}

// Count returns the number of available forms, and an occurring error
func (j *JSONSource) Count() (int64, error) {
	return int64(len(j.forms0)), nil
}
