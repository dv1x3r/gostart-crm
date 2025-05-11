package productstatus

import "errors"

var (
	ErrProductStatusExists         = errors.New("already exists")
	ErrProductStatusHasRelated     = errors.New("has related objects")
	ErrProductStatusInvalidReorder = errors.New("invalid reorder parameters")
)
