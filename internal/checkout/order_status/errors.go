package orderstatus

import "errors"

var (
	ErrOrderStatusExists         = errors.New("already exists")
	ErrOrderStatusHasRelated     = errors.New("has related objects")
	ErrOrderStatusInvalidReorder = errors.New("invalid reorder parameters")
)
