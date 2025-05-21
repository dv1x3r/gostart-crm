package supplier

import (
	"github.com/dv1x3r/w2go/w2"
	"github.com/gosimple/slug"
)

type Supplier struct {
	ID          int                 `json:"id"`
	Code        w2.Editable[string] `json:"code"`
	Name        w2.Editable[string] `json:"name"`
	Description w2.Editable[string] `json:"description"`
	IsPublished w2.Editable[bool]   `json:"is_published"`

	RelatedProducts   int `json:"related_products"`
	PublishedProducts int `json:"published_products"`
}

func (s *Supplier) Slugify() string {
	if s.Code.V == "" {
		return ""
	}
	return slug.Make(s.Code.V)
}

func (s *Supplier) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&s.ID,
		&s.Code,
		&s.Name,
		&s.Description,
		&s.IsPublished,
		&s.RelatedProducts,
		&s.PublishedProducts,
	)
}
