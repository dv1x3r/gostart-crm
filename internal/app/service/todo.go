package service

import "w2go/internal/app/model"

type Todo struct {
	data []model.TodoDTO
}

func NewTodo() *Todo {
	data := []model.TodoDTO{
		{ID: 0, Name: "Demo", Description: new(string)},
		{ID: 1, Name: "List", Description: new(string)},
		{ID: 2, Name: "End", Quantity: new(int64)},
	}
	*data[0].Description = "<b>Description</b>"
	*data[1].Description = "Placeholder field<script>alert('XSS Test')</script>"
	*data[2].Quantity = 5
	return &Todo{data: data}
}

func (ts *Todo) GetTodoW2Grid(req model.W2GridDataRequest) (model.TodoW2GridResponse, error) {
	return model.TodoW2GridResponse{Status: "success", Records: ts.data}, nil
}

func (ts *Todo) DeleteTodoW2Action(idList []int64) (model.W2BaseResponse, error) {
	newData := []model.TodoDTO{}

out:
	for _, td := range ts.data {
		for _, deleteID := range idList {
			if td.ID == deleteID {
				continue out
			}
		}
		newData = append(newData, td)
	}

	ts.data = newData
	return model.W2BaseResponse{Status: "success"}, nil
}

func (ts *Todo) PatchTodoW2Action(t []model.TodoPatchDTO) (model.W2BaseResponse, error) {
	return model.W2BaseResponse{Status: "success"}, nil
}

func (ts *Todo) UpsertTodoW2Form(id int64, t model.TodoDTO) (model.W2BaseResponse, error) {
	if id == 0 {
		ts.data = append(ts.data, t)
	}
	return model.W2BaseResponse{Status: "success"}, nil
}
