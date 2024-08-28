package model

type OrderLine struct {
	ID       int64   `json:"id" db:"id"`
	Code     string  `json:"code" db:"code"`
	Product  string  `json:"product" db:"product"`
	Quantity float64 `json:"quantity" db:"quantity"`
	Price    float64 `json:"price" db:"price"`
	Total    float64 `json:"total" db:"total"`
	Summary  float64 `json:"summary" db:"summary"`
}

type OrderLineW2GridResponse = W2GridDataResponse[OrderLine, any]

type OrderLineSummary struct {
	Total float64 `json:"total"`
}
