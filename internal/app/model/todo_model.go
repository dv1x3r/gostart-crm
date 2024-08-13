package model

import "encoding/json"

type TodoDTO struct {
	ID          int64   `json:"id" db:"id" validate:"number"`
	Name        string  `json:"name" db:"name" validate:"required,max=8"`
	Description *string `json:"description,omitempty" db:"description" validate:"required,max=16"`
	Quantity    *int64  `json:"quantity,omitempty" db:"quantity" validate:"number,min=-2147483647,max=2147483647"`

	FieldMap FieldMap
}

func (t *TodoDTO) GetValidateFieldMap() FieldMap {
	return t.FieldMap
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
	t.ID = getValue[int64]("id", "ID", raw, t.FieldMap)
	t.Name = getValue[string]("name", "Name", raw, t.FieldMap)
	t.Description = getValuePtr[string]("description", "Description", raw, t.FieldMap)
	t.Quantity = getValuePtr[int64]("quantity", "Quantity", raw, t.FieldMap)

	return nil
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2SaveRequest = W2GridSaveRequest[TodoDTO]
