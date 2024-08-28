package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type AttributeValue struct {
	ID              int64  `json:"id" db:"id"`
	AttributeSetID  int64  `json:"attribute_set_id" db:"attribute_set_id"`
	Name            string `json:"name" db:"name" validate:"required,max=128"`
	RelatedProducts int64  `json:"related_products" db:"related_products"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *AttributeValue) UnmarshalJSON(data []byte) error {
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

type AttributeValueW2GridResponse = W2GridDataResponse[AttributeValue, any]
type AttributeValueW2SaveRequest = W2GridSaveRequest[AttributeValue]
