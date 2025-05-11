package category

import "errors"

var (
	ErrCategoryExists          = errors.New("already exists")
	ErrCategoryNotFound        = errors.New("not found")
	ErrCategoryHasRelated      = errors.New("has related objects")
	ErrCategoryInvalidReorder  = errors.New("invalid reorder parameters")
	ErrCategoryInvalidCircular = errors.New("invalid circular tree")
)
