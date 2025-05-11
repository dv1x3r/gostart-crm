package supplier

import "errors"

var (
	ErrSupplierExists         = errors.New("already exists")
	ErrSupplierHasRelated     = errors.New("has related objects")
	ErrSupplierInvalidReorder = errors.New("invalid reorder parameters")
)
