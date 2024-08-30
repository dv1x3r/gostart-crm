package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type OrderStatus struct {
	ID            int64  `json:"id" db:"id" validate:"number"`
	Name          string `json:"name" db:"name" validate:"required,max=16"`
	Color         string `json:"color" db:"color" validate:"required,max=16"`
	InCounter     bool   `json:"in_counter" db:"in_counter"`
	RelatedOrders int64  `json:"related_orders" db:"related_orders"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *OrderStatus) UnmarshalJSON(data []byte) error {
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
	t.Color = getValue[string](raw, "color", t.Partial, "Color")
	t.InCounter = getValue[bool](raw, "in_counter", t.Partial, "InCounter")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type OrderStatusW2GridResponse = W2GridDataResponse[OrderStatus, any]
type OrderStatusW2SaveRequest = W2GridSaveRequest[OrderStatus]
