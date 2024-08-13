package model

import (
	"encoding/json"
)

type TodoDTO struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Quantity    *int64  `json:"quantity,omitempty" db:"quantity"`

	FieldMap FieldMap `json:"-" db:"-"`
}

func (t *TodoDTO) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(FieldRaw)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.FieldMap = make(FieldMap)
	t.ID = getValue[int64]("id", &raw, "ID", &t.FieldMap)

	return nil
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2SaveRequest = W2GridSaveRequest[TodoDTO]
type TodoW2FormResponse = W2FormResponse[TodoDTO]
type TodoW2FormRequest = W2FormRequest[TodoDTO]
