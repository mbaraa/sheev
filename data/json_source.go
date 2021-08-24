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
	forms0  []models.Form
	fields0 []models.Field

	// maps, much speed wow!
	forms  map[string]models.Form    // formName -> form
	fields map[string][]models.Field // formName -> fields
}

// NewJSONSource returns a new JSONSource instance
func NewJSONSource(jsonDir string) (j *JSONSource) {
	return new(JSONSource).
		loadJSONFiles(jsonDir).
		initDataMaps()
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

func (j *JSONSource) initDataMaps() *JSONSource {
	// malloc
	j.forms = make(map[string]models.Form)
	j.fields = make(map[string][]models.Field)

	for _, field := range j.fields0 {
		j.fields[field.FormName] = append(j.fields[field.FormName], field)
	}

	for formIndex, form := range j.forms0 {

		j.forms[form.Name] = models.Form{
			Name:       form.Name,
			Fields:     j.fields[form.Name], // the whole new form struct is just for this little fucker 🙂
			B64FormImg: form.B64FormImg,
		}

		j.forms0[formIndex].Fields = j.fields[form.Name]
	}

	return j
}

// ExistsByName reports whether the FormGenerator exists or not, and an occurring error
func (j *JSONSource) ExistsByName(name string) bool {
	_, formExists := j.forms[name]
	return formExists
}

// Get returns a form depending on its name, and an occurring error
func (j *JSONSource) Get(name string) (models.Form, error) {
	if form, formExists := j.forms[name]; formExists {
		return form, nil
	}
	return models.Form{}, errors.ErrNoFormFound
}

// GetAll returns all available forms, and an occurring error
func (j *JSONSource) GetAll() ([]models.Form, error) {
	return j.forms0, nil
}

// Count returns the number of available forms, and an occurring error
func (j *JSONSource) Count() (int64, error) {
	return int64(len(j.forms0)), nil
}
