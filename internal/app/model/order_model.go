package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type OrderStatusEmbed struct {
	StatusID    *int64  `json:"id" db:"order_status_id"`
	StatusName  *string `json:"text" db:"order_status_name"`
	StatusColor *string `json:"color" db:"order_status_color"`
}

type PaymentMethodEmbed struct {
	PaymentMethodID   *int64  `json:"id" db:"payment_method_id"`
	PaymentMethodName *string `json:"text" db:"payment_method_name"`
}

type Order struct {
	ID              int64   `json:"id" db:"id"`
	Email           string  `json:"email" db:"email" validate:"required,max=256"`
	FirstName       *string `json:"first_name" db:"first_name" validate:"required,max=256"`
	LastName        *string `json:"last_name" db:"last_name" validate:"required,max=256"`
	PhoneNumber     *string `json:"phone_number" db:"phone_number" validate:"required,max=256"`
	DeliveryAddress *string `json:"delivery_address" db:"delivery_address" validate:"omitempty,max=32768"`
	Comment         *string `json:"comment" db:"comment" validate:"omitempty,max=32768"`
	Language        *string `json:"language" db:"language" validate:"omitempty,max=32768"`
	Notes           *string `json:"notes" db:"notes" validate:"omitempty,max=32768"`
	FullName        *string `json:"full_name" db:"full_name"`
	CreatedAt       string  `json:"created_at" db:"created_at"`
	UpdatedAt       string  `json:"updated_at" db:"updated_at"`
	Total           float64 `json:"total" db:"total"`

	OrderStatusEmbed   `json:"status"`
	PaymentMethodEmbed `json:"payment"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *Order) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Partial = make(map[string]struct{})
	t.ID = getValue[int64](raw, "id", t.Partial, "ID")
	t.Email = getValue[string](raw, "email", t.Partial, "Email")
	t.FirstName = getValuePtr[string](raw, "first_name", t.Partial, "FirstName")
	t.LastName = getValuePtr[string](raw, "last_name", t.Partial, "LastName")
	t.PhoneNumber = getValuePtr[string](raw, "phone_number", t.Partial, "PhoneNumber")
	t.Language = getValuePtr[string](raw, "language", t.Partial, "Language")
	t.Notes = getValuePtr[string](raw, "notes", t.Partial, "Notes")
	t.OrderStatusEmbed = getValue[OrderStatusEmbed](raw, "status", t.Partial, "OrderStatusEmbed.StatusID")
	t.PaymentMethodEmbed = getValue[PaymentMethodEmbed](raw, "payment", t.Partial, "PaymentMethodEmbed.PaymentMethodID")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type OrderW2GridResponse = W2GridDataResponse[Order, any]
type OrderW2FormRequest = W2FormRequest[Order]
type OrderW2FormResponse = W2FormResponse[Order]
