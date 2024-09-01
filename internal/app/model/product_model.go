package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"

	"github.com/gosimple/slug"
)

type ProductCategoryEmbed struct {
	CategoryID               int64  `json:"id" db:"category_id"`
	CategoryHierarchy        string `json:"text" db:"category_hierarchy"`
	CategoryAttributeGroupID *int64 `json:"attribute_group_id" db:"category_attribute_group_id"`
}

type ProductSupplierEmbed struct {
	SupplierID   int64  `json:"id" db:"supplier_id"`
	SupplierName string `json:"text" db:"supplier_name"`
}

type ProductBrandEmbed struct {
	BrandID   int64  `json:"id" db:"brand_id"`
	BrandName string `json:"text" db:"brand_name"`
}

type ProductStatusEmbed struct {
	StatusID    *int64  `json:"id" db:"status_id"`
	StatusName  *string `json:"text" db:"status_name"`
	StatusColor *string `json:"color" db:"status_color"`
}

type Product struct {
	ID          int64    `json:"id" db:"id" validate:"number"`
	Code        string   `json:"code" db:"code" validate:"required,max=32"`
	Name        string   `json:"name" db:"name" validate:"required,max=1024"`
	Description *string  `json:"description" db:"description" validate:"omitempty,max=32768"`
	Quantity    float64  `json:"quantity" db:"quantity" validate:"number,min=-2147483647,max=2147483647"`
	Price       *float64 `json:"price" db:"price" validate:"omitempty,number,min=-2147483647,max=2147483647"`
	IsPublished bool     `json:"is_published" db:"is_published"`
	CreatedAt   string   `json:"created_at" db:"created_at"`
	UpdatedAt   string   `json:"updated_at" db:"updated_at"`
	IsAvailable bool     `json:"is_available" db:"is_available"`
	ProductURL  string   `json:"product_url" db:"product_url"`

	ProductCategoryEmbed `json:"category"`
	ProductSupplierEmbed `json:"supplier"`
	ProductBrandEmbed    `json:"brand"`
	ProductStatusEmbed   `json:"status"`

	Partial    map[string]struct{} `json:"-" db:"-"`
	Attributes []ProductAttribute  `json:"attributes,omitempty" db:"-"`
}

func (dto *Product) Slugify() string {
	if dto.Code == "" || dto.Name == "" {
		return ""
	}

	s := slug.Make(dto.Code + "-" + dto.Name)
	if len(s) > 64 {
		return s[:64]
	}

	return s
}

func (t *Product) UnmarshalJSON(data []byte) error {
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
	t.Quantity = getValue[float64](raw, "quantity", t.Partial, "Quantity")
	t.Price = getValuePtr[float64](raw, "price", t.Partial, "Price")
	t.IsPublished = getValue[bool](raw, "is_published", t.Partial, "IsPublished")
	t.ProductCategoryEmbed = getValue[ProductCategoryEmbed](raw, "category", t.Partial, "ProductCategoryEmbed.CategoryID")
	t.ProductSupplierEmbed = getValue[ProductSupplierEmbed](raw, "supplier", t.Partial, "ProductSupplierEmbed.SupplierID")
	t.ProductBrandEmbed = getValue[ProductBrandEmbed](raw, "brand", t.Partial, "ProductBrandEmbed.BrandID")
	t.ProductStatusEmbed = getValue[ProductStatusEmbed](raw, "status", t.Partial, "ProductStatusEmbed.StatusID")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type ProductW2GridResponse = W2GridDataResponse[Product, any]
type ProductW2SaveRequest = W2GridSaveRequest[Product]
type ProductW2FormResponse = W2FormResponse[Product]
type ProductW2FormRequest = W2FormRequest[Product]
