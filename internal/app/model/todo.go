package model

type TodoDTO struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Quantity    *int64  `json:"quantity,omitempty"`
}

type TodoPatchDTO struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Quantity    *any    `json:"quantity,omitempty"`
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2PatchRequest = W2GridPatchRequest[TodoPatchDTO]
type TodoW2FormRequest = W2FormRequest[TodoDTO]
type TodoW2FormResponse = W2FormResponse[TodoDTO]
