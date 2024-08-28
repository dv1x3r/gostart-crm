package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type ProductStatusDropdown struct {
	W2Dropdown
	Color string `json:"color" db:"color"`
}

type ProductStatus struct {
	ID              int64  `json:"id" db:"id" validate:"number"`
	Name            string `json:"name" db:"name" validate:"required,max=16"`
	Color           string `json:"color" db:"color" validate:"required,max=16"`
	RelatedProducts int64  `json:"related_products" db:"related_products"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *ProductStatus) UnmarshalJSON(data []byte) error {
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

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type ProductStatusW2GridResponse = W2GridDataResponse[ProductStatus, any]
type ProductStatusW2SaveRequest = W2GridSaveRequest[ProductStatus]
