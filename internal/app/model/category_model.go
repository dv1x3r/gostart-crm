package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"

	"github.com/gosimple/slug"
)

type CategoryDropdown struct {
	W2Dropdown
	AttributeGroupID *int64 `json:"attribute_group_id" db:"attribute_group_id"`
}

type CategoryParentEmbed struct {
	ParentID   *int64  `json:"id" db:"parent_id"`
	ParentText *string `json:"text" db:"parent_text"`
}

type CategoryAttributeGroupEmbed struct {
	AttributeGroupID   *int64  `json:"id" db:"attribute_group_id"`
	AttributeGroupText *string `json:"text" db:"attribute_group_text"`
}

type Category struct {
	ID              int64   `json:"id" db:"id" validate:"number"`
	Name            string  `json:"name" db:"name" validate:"required,max=64"`
	Icon            *string `json:"icon" db:"icon" validate:"omitempty,max=128"`
	IsPublished     bool    `json:"is_published" db:"is_published"`
	RelatedProducts int64   `json:"related_products" db:"related_products"`
	Hierarchy       string  `json:"hierarchy" db:"hierarchy"`
	CategoryURL     string  `json:"-" db:"category_url"`

	CategoryParentEmbed         `json:"parent"`
	CategoryAttributeGroupEmbed `json:"attribute_group"`

	Partial  map[string]struct{} `json:"-" db:"-"`
	Children []Category          `json:"children" db:"-"`
}

func (dto *Category) Slugify() string {
	if dto.Name == "" {
		return ""
	}
	return slug.Make(dto.Name)
}

func (t *Category) UnmarshalJSON(data []byte) error {
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
	t.Icon = getValuePtr[string](raw, "icon", t.Partial, "Icon")
	t.IsPublished = getValue[bool](raw, "is_published", t.Partial, "IsPublished")
	t.CategoryParentEmbed = getValue[CategoryParentEmbed](raw, "parent", t.Partial, "CategoryParentEmbed.ParentID")
	t.CategoryAttributeGroupEmbed = getValue[CategoryAttributeGroupEmbed](raw, "attribute_group", t.Partial, "CategoryAttributeGroupEmbed.AttributeGroupID")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type CategoryW2GridResponse = W2GridDataResponse[Category, any]
type CategoryW2SaveRequest = W2GridSaveRequest[Category]
