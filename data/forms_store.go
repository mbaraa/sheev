package data

import (
	"github.com/mbaraa/sheev/models"
)

type FormsGetterRepo interface {
	// ExistsByName reports whether the FormGenerator exists or not, and an occurring error
	ExistsByName(string) bool
	// Get returns a form depending on its name, and an occurring error
	Get(string) (models.Form, error)
	// GetAll returns all available forms, and an occurring error
	GetAll() ([]models.Form, error)
	// Count returns the number of available forms, and an occurring error
	Count() (int64, error)
}

// FormsStore represents a general forms data source
type FormsStore interface {
	FormsGetterRepo
}
