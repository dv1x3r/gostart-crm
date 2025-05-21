package validator

import (
	ut "github.com/go-playground/universal-translator"
	playgroundValidator "github.com/go-playground/validator/v10"
)

var validator *Validator

type Validator struct {
	validate    *playgroundValidator.Validate
	translators map[string]ut.Translator
}

func init() {
	validator = &Validator{
		validate:    playgroundValidator.New(playgroundValidator.WithRequiredStructEnabled()),
		translators: map[string]ut.Translator{},
	}
}

func GetValidator() *Validator {
	return validator
}

func (v *Validator) Validate(s interface{}) error {
	return v.validate.Struct(s)
}
