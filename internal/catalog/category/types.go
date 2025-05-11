package category

import (
	"github.com/dv1x3r/w2go/w2"
	"github.com/gosimple/slug"
)

type Category struct {
	ID             int                 `json:"id"`
	Name           w2.Editable[string] `json:"name"`
	Icon           w2.Editable[string] `json:"icon"`
	IsPublished    w2.Editable[bool]   `json:"is_published"`
	Parent         w2.EditableDropdown `json:"parent"`
	AttributeGroup w2.EditableDropdown `json:"attribute_group"`

	RelatedProducts int    `json:"related_products"`
	Hierarchy       string `json:"hierarchy"`
	CategoryURL     string `json:"category_url"`

	Children []Category `json:"children"`
}

func (c *Category) Slugify() string {
	if c.Name.V == "" {
		return ""
	}
	return slug.Make(c.Name.V)
}

func (c *Category) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&c.ID,
		&c.Name,
		&c.Icon,
		&c.IsPublished,
		&c.Parent.ID,
		&c.Parent.Text,
		&c.AttributeGroup.ID,
		&c.AttributeGroup.Text,
		&c.Hierarchy,
		&c.CategoryURL,
		&c.RelatedProducts,
	)
}
