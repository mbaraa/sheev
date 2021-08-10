package errors

import "errors"

var (
	ErrNotImplemented       = errors.New("not implemented")
	ErrFieldOverflowsParent = errors.New("field overflows parent")
	ErrNoFormFound          = errors.New("no matching forms were found")
	ErrNoFieldFound         = errors.New("no such field found")
)
