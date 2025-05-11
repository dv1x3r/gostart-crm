package orderstatus

import "github.com/dv1x3r/w2go/w2"

type OrderStatus struct {
	ID            int                 `json:"id"`
	Name          w2.Editable[string] `json:"name"`
	Color         w2.Editable[string] `json:"color"`
	RelatedOrders int                 `json:"related_orders"`
}

func (s *OrderStatus) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&s.ID,
		&s.Name,
		&s.Color,
		&s.RelatedOrders,
	)
}
