package paymentmethod

import "github.com/dv1x3r/w2go/w2"

type PaymentMethod struct {
	ID            int                 `json:"id"`
	Name          w2.Editable[string] `json:"name"`
	RelatedOrders int                 `json:"related_orders"`
}

func (m *PaymentMethod) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&m.ID,
		&m.Name,
		&m.RelatedOrders,
	)
}
