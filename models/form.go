package models

// Form represents a form :]
type Form struct {
	Name       string   `json:"name"`
	Fields     []*Field `json:"fields"`
	B64FormImg string   `json:"form_img"`
}

func (f *Form) CopyForm() Form {
	return Form{
		Name:       f.Name,
		Fields:     f.Fields,
		B64FormImg: f.B64FormImg,
	}
}
