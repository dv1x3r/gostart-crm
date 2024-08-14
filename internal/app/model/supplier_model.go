package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"

	"github.com/gosimple/slug"
)

type Supplier struct {
	ID              int64   `json:"id" db:"id" validate:"number"`
	Slug            string  `json:"-" db:"slug"`
	Code            string  `json:"code" db:"code" validate:"required,max=8"`
	Name            string  `json:"name" db:"name" validate:"required,max=16"`
	Description     *string `json:"description" db:"description" validate:"omitempty,max=256"`
	IsPublished     bool    `json:"is_published" db:"is_published"`
	RelatedProducts int64   `json:"related_products" db:"related_products"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (dto *Supplier) Slugify() string {
	if dto.Code == "" {
		return ""
	}
	return slug.Make(dto.Code)
}

func (t *Supplier) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Partial = make(map[string]struct{})
	t.ID = getValue[int64](raw, "id", t.Partial, "ID")
	t.Code = getValue[string](raw, "code", t.Partial, "Code")
	t.Name = getValue[string](raw, "name", t.Partial, "Name")
	t.Description = getValuePtr[string](raw, "description", t.Partial, "Description")
	t.IsPublished = getValue[bool](raw, "is_published", t.Partial, "IsPublished")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type SupplierW2GridResponse = W2GridDataResponse[Supplier, any]
type SupplierW2SaveRequest = W2GridSaveRequest[Supplier]
