package product

import (
	"github.com/dv1x3r/w2go/w2"
	"github.com/gosimple/slug"
)

type ListDataRequest struct {
	Limit    int
	Offset   int
	Search   string
	SearchBy string
	Filters  string
}

type Product struct {
	ID          int                  `json:"id"`
	Code        w2.Editable[string]  `json:"code"`
	Name        w2.Editable[string]  `json:"name"`
	Description w2.Editable[string]  `json:"description"`
	Quantity    w2.Editable[float64] `json:"quantity"`
	Price       w2.Editable[float64] `json:"price"`
	IsPublished w2.Editable[bool]    `json:"is_published"`

	Category w2.EditableDropdown `json:"category"`
	Supplier w2.EditableDropdown `json:"supplier"`
	Brand    w2.EditableDropdown `json:"brand"`
	Status   struct {
		w2.EditableDropdown
		Color w2.Editable[string] `json:"color"`
	} `json:"status"`

	ProductURL string `json:"product_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`

	Attributes []ProductAttribute `json:"attributes,omitempty"`
}

func (p *Product) Slugify() string {
	if p.Code.V == "" || p.Name.V == "" {
		return ""
	}

	s := slug.Make(p.Code.V + "-" + p.Name.V)
	if len(s) > 64 {
		return s[:64]
	}

	return s
}

func (p *Product) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&p.ID,
		&p.Code,
		&p.Name,
		&p.Description,
		&p.Quantity,
		&p.Price,
		&p.IsPublished,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Category.ID,
		&p.Category.Text,
		&p.Supplier.ID,
		&p.Supplier.Text,
		&p.Brand.ID,
		&p.Brand.Text,
		&p.Status.ID,
		&p.Status.Text,
		&p.Status.Color,
		&p.ProductURL,
	)
}

type ProductAttribute struct {
	AttributeSetID   int                 `json:"id"`
	AttributeSetName string              `json:"name"`
	AttributeValue   w2.EditableDropdown `json:"value"`
}

func (a *ProductAttribute) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&a.AttributeSetID,
		&a.AttributeSetName,
		&a.AttributeValue.ID,
		&a.AttributeValue.Text,
	)
}
