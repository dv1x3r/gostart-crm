package storage

import "errors"

var (
	ErrConstraintUnique     = errors.New("storage: unique constraint failed")
	ErrConstraintTrigger    = errors.New("storage: constraint trigger failed")
	ErrConstraintForeignKey = errors.New("storage: constraint foreign key failed")
	ErrInvalidTreeCircular  = errors.New("storage: circular tree update failed")
)
