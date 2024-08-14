package model

import (
	"encoding/json"
	"gostart-crm/internal/app/utils"
)

type TodoDTO struct {
	ID          int64   `json:"id" db:"id" validate:"number"`
	Name        string  `json:"name" db:"name" validate:"required,max=8"`
	Description *string `json:"description,omitempty" db:"description" validate:"required,max=16"`
	Quantity    *int64  `json:"quantity,omitempty" db:"quantity" validate:"number,min=-2147483647,max=2147483647"`

	Partial map[string]struct{} `json:"-" db:"-"`
}

func (t *TodoDTO) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Partial = make(map[string]struct{})
	t.ID = getValue[int64](raw, "id", t.Partial, "ID")
	t.Name = getValue[string](raw, "name", t.Partial, "Name")
	t.Description = getValuePtr[string](raw, "description", t.Partial, "Description")
	t.Quantity = getValuePtr[int64](raw, "quantity", t.Partial, "Quantity")

	return utils.GetValidator().ValidatePartial(t, t.Partial)
}

type TodoW2GridResponse = W2GridDataResponse[TodoDTO, any]
type TodoW2SaveRequest = W2GridSaveRequest[TodoDTO]
