package productstatus

import "github.com/dv1x3r/w2go/w2"

type ProductStatus struct {
	ID    int                 `json:"id"`
	Name  w2.Editable[string] `json:"name"`
	Color w2.Editable[string] `json:"color"`

	RelatedProducts int `json:"related_products"`
}

func (s *ProductStatus) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&s.ID,
		&s.Name,
		&s.Color,
		&s.RelatedProducts,
	)
}

type ProductStatusDropdown struct {
	w2.DropdownValue
	Color string `json:"color"`
}
