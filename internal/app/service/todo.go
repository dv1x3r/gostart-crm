package service

import (
	"w2go/internal/app/model"
)

type TodoStorager interface {
	SelectList(model.QuerySelectList) ([]model.TodoFromDB, error)
	// Insert(model.TodoDTO) (int64, error)
	// Update(int64, model.TodoDTO) error
	// Patch(int64, model.TodoPatchDTO) error
	// Delete(int64) error
}

type Todo struct {
	st   TodoStorager
	data []model.TodoDTO
}

func NewTodo(st TodoStorager) *Todo {
	data := []model.TodoDTO{
		{ID: 0, Name: "Demo", Description: new(string)},
		{ID: 1, Name: "List", Description: new(string)},
		{ID: 2, Name: "End", Quantity: new(int64)},
	}
	*data[0].Description = "<b>Description</b>"
	*data[1].Description = "Placeholder field<script>alert('XSS Test')</script>"
	*data[2].Quantity = 5
	return &Todo{st: st, data: data}
}

func (ts *Todo) GetTodoW2Grid(req model.W2GridDataRequest) (model.TodoW2GridResponse, error) {
	rows, err := ts.st.SelectList(req.ToQuerySelectList())
	todos := make([]model.TodoDTO, len(rows))
	var count int64

	for i, row := range rows {
		count = row.Count
		todos[i] = row.TodoDTO
	}

	if err != nil {
		return model.TodoW2GridResponse{Status: "error"}, err
	} else {
		return model.TodoW2GridResponse{Status: "success", Records: todos, Total: count}, nil
	}
}

func (ts *Todo) DeleteTodoW2Action(req model.W2GridDeleteRequest) (model.W2BaseResponse, error) {
	newData := []model.TodoDTO{}

out:
	for _, td := range ts.data {
		for _, deleteID := range req.ID {
			if td.ID == deleteID {
				continue out
			}
		}
		newData = append(newData, td)
	}

	ts.data = newData
	return model.W2BaseResponse{Status: "success"}, nil
}

func (ts *Todo) PatchTodoW2Action(req model.TodoW2PatchRequest) (model.W2BaseResponse, error) {
	return model.W2BaseResponse{Status: "error", Message: "not implemented"}, nil
}

func (ts *Todo) GetTodoW2Form(req model.TodoW2FormRequest) (model.TodoW2FormResponse, error) {
	for _, row := range ts.data {
		if row.ID == req.RecID {
			return model.TodoW2FormResponse{Status: "success", Record: row}, nil
		}
	}
	return model.TodoW2FormResponse{Status: "error", Message: "not found"}, nil
}

func (ts *Todo) UpsertTodoW2Form(req model.TodoW2FormRequest) (model.W2BaseResponse, error) {
	if req.RecID == 0 {
		ts.data = append(ts.data, req.Record)
	} else {
		for i, row := range ts.data {
			if row.ID == req.RecID {
				ts.data[i] = req.Record
				break
			}
		}
	}
	return model.W2BaseResponse{Status: "success"}, nil
}
