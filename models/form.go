package models

// Form represents a form :]
type Form struct {
	Name       string  `json:"name"`
	Fields     []Field `json:"fields"`
	B64FormImg string  `json:"form_img"`
}

// CopyForm returns a copy of the current form (to avoid pointers fuckery)
func (f *Form) CopyForm() Form {
	newFields := make([]Field, len(f.Fields))
	copy(newFields, f.Fields)

	return Form{
		Name:       f.Name,
		Fields:     newFields,
		B64FormImg: f.B64FormImg,
	}
}
