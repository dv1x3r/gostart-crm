package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type ProductAttributeValueEmbed struct {
	AttributeValueID *int64  `json:"id" db:"value_id" validate:"required"`
	AttributeValue   *string `json:"text" db:"value"`
}

type ProductAttribute struct {
	AttributeSetID   int64  `json:"id" db:"id"`
	AttributeSetName string `json:"name" db:"name"`

	ProductAttributeValueEmbed `json:"value"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *ProductAttribute) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Partial = make(map[string]struct{})
	t.AttributeSetID = getValue[int64](raw, "id", t.Partial, "AttributeSetID")
	t.ProductAttributeValueEmbed = getValue[ProductAttributeValueEmbed](raw, "value", t.Partial, "ProductAttributeValueEmbed.AttributeValueID")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type ProductAttributeW2GridResponse = W2GridDataResponse[ProductAttribute, any]
type ProductAttributeW2SaveRequest = W2GridSaveRequest[ProductAttribute]
