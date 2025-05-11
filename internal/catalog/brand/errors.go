package brand

import "errors"

var (
	ErrBrandExists     = errors.New("already exists")
	ErrBrandHasRelated = errors.New("has related objects")
)
