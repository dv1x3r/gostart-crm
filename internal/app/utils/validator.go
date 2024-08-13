package utils

import (
	"maps"
	"slices"
	"sync"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validate *validator.Validate
}

func (ev *CustomValidator) Validate(s interface{}) error {
	return ev.validate.Struct(s)
}

func (ev *CustomValidator) ValidatePartial(s interface{}, partials map[string]struct{}) error {
	return ev.validate.StructPartial(s, slices.Collect(maps.Keys(partials))...)
}

var (
	validateOnce    sync.Once
	customValidator *CustomValidator
)

func GetValidator() *CustomValidator {
	validateOnce.Do(func() {
		validate := validator.New(validator.WithRequiredStructEnabled())
		customValidator = &CustomValidator{
			validate: validate,
		}
	})

	return customValidator
}
