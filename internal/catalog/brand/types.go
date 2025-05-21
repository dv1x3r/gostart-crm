package brand

import "github.com/dv1x3r/w2go/w2"

type Brand struct {
	ID   int                 `json:"id"`
	Name w2.Editable[string] `json:"name"`

	RelatedProducts   int `json:"related_products"`
	PublishedProducts int `json:"published_products"`
}

func (b *Brand) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&b.ID,
		&b.Name,
		&b.RelatedProducts,
		&b.PublishedProducts,
	)
}
