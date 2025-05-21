package paymentmethod

import "errors"

var (
	ErrPaymentMethodExists         = errors.New("already exists")
	ErrPaymentMethodHasRelated     = errors.New("has related objects")
	ErrPaymentMethodInvalidReorder = errors.New("invalid reorder parameters")
)
