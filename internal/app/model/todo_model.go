package model

type TodoDTO struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Quantity    *int64  `json:"quantity,omitempty" db:"quantity"`
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2SaveRequest = W2GridSaveRequest[TodoDTO]
type TodoW2FormResponse = W2FormResponse[TodoDTO]
type TodoW2FormRequest = W2FormRequest[TodoDTO]
