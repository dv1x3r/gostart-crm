package model

type TodoDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type TodoW2Grid = W2GridDataResponse[TodoDTO, any]
type TodoW2Form = W2FormSubmit[TodoDTO, int64]
