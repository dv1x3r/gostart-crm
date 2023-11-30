package model

type TodoDTO struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Quantity    *int64  `json:"quantity,omitempty" db:"quantity"`
}

type TodoFromDB struct {
	TodoDTO
	Count int64 `db:"count"`
}

type TodoPartialDTO struct {
	ID       int64 `json:"id"`
	Quantity *any  `json:"quantity"`
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2PatchRequest = W2GridPatchRequest[TodoPartialDTO]
type TodoW2FormRequest = W2FormRequest[TodoDTO]
type TodoW2FormResponse = W2FormResponse[TodoDTO]
