package product

import "errors"

var (
	ErrProductExists              = errors.New("already exists")
	ErrProductNotFound            = errors.New("not found")
	ErrProductDeleteProtect       = errors.New("delete protection rows limit")
	ErrProductMediaNotFound       = errors.New("media not found")
	ErrProductMediaInvalidReorder = errors.New("invalid media reorder parameters")
	ErrProductAttributeNotFound   = errors.New("attribute not found")
)
