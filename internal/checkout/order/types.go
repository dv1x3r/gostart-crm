package order

import "github.com/dv1x3r/w2go/w2"

type Order struct {
	ID int `json:"id" db:"id"`

	Email       w2.Editable[string] `json:"email"`
	FirstName   w2.Editable[string] `json:"first_name"`
	LastName    w2.Editable[string] `json:"last_name"`
	PhoneNumber w2.Editable[string] `json:"phone_number"`

	Status        w2.EditableDropdown `json:"status"`
	PaymentMethod w2.EditableDropdown `json:"payment"`

	Color           w2.Editable[string] `json:"color"`
	DeliveryAddress w2.Editable[string] `json:"delivery_address"`
	Comment         w2.Editable[string] `json:"comment"`

	Total     float64 `json:"total"`
	FullName  string  `json:"full_name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`

	Lines []OrderLine `json:"lines"`
}

func (o *Order) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&o.ID,
		&o.Email,
		&o.FullName,
		&o.FirstName,
		&o.LastName,
		&o.PhoneNumber,
		&o.DeliveryAddress,
		&o.Comment,
		&o.CreatedAt,
		&o.UpdatedAt,
		&o.Status.ID,
		&o.Status.Text,
		&o.Color,
		&o.PaymentMethod.ID,
		&o.PaymentMethod.Text,
		&o.Total,
	)
}

type OrderLine struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Product  string  `json:"product"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}

func (l *OrderLine) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&l.ID,
		&l.Code,
		&l.Product,
		&l.Quantity,
		&l.Price,
		&l.Total,
	)
}
