package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type PaymentMethod struct {
	ID            int64  `json:"id" db:"id" validate:"number"`
	Name          string `json:"name" db:"name" validate:"required,max=32"`
	RelatedOrders int64  `json:"related_orders" db:"related_orders"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *PaymentMethod) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Partial = make(map[string]struct{})
	t.ID = getValue[int64](raw, "id", t.Partial, "ID")
	t.Name = getValue[string](raw, "name", t.Partial, "Name")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type PaymentMethodW2GridResponse = W2GridDataResponse[PaymentMethod, any]
type PaymentMethodW2SaveRequest = W2GridSaveRequest[PaymentMethod]
