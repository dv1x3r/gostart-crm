package order

import "errors"

var (
	ErrOrderDeleteProtect = errors.New("delete protection rows limit")
)
