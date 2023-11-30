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

type TodoPatchDTO struct {
	ID          int64 `json:"id"`
	Name        *any  `json:"name,omitempty"`
	Description *any  `json:"description,omitempty"`
	Quantity    *any  `json:"quantity,omitempty"`
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2PatchRequest = W2GridPatchRequest[TodoPatchDTO]
type TodoW2FormRequest = W2FormRequest[TodoDTO]
type TodoW2FormResponse = W2FormResponse[TodoDTO]
