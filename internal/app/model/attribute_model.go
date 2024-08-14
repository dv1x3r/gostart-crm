package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type AttributeGroup struct {
	ID   int64  `json:"id" db:"id" validate:"number"`
	Name string `json:"name" db:"name" validate:"required,max=128"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *AttributeGroup) UnmarshalJSON(data []byte) error {
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

type AttributeGroupW2GridResponse = W2GridDataResponse[AttributeGroup, any]

type AttributeSet struct {
	ID               int64  `json:"id" db:"id" validate:"number"`
	AttributeGroupID int64  `json:"attribute_group_id" db:"attribute_group_id"`
	Name             string `json:"name" db:"name" validate:"required,max=128"`
	InBox            bool   `json:"in_box" db:"in_box"`
	InFilter         bool   `json:"in_filter" db:"in_filter"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *AttributeSet) UnmarshalJSON(data []byte) error {
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
	t.InBox = getValue[bool](raw, "in_box", t.Partial, "InBox")
	t.InFilter = getValue[bool](raw, "in_filter", t.Partial, "InFilter")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type AttributeSetW2GridResponse = W2GridDataResponse[AttributeSet, any]
type AttributeSetW2SaveRequest = W2GridSaveRequest[AttributeSet]

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
