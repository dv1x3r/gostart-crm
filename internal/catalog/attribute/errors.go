package attribute

import "errors"

var (
	ErrAttributeExists         = errors.New("already exists")
	ErrAttributeNotFound       = errors.New("not found")
	ErrAttributeHasRelated     = errors.New("has related objects")
	ErrAttributeInvalidReorder = errors.New("invalid reorder parameters")
)
