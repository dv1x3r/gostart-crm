package service

import "w2go/internal/app/model"

type Todo struct {
}

func (ts *Todo) GetTodoW2Grid(req model.W2GridDataRequest) (model.TodoW2Grid, error) {
	data := []model.TodoDTO{
		{ID: 0, Name: "Demo", Description: new(string)},
		{ID: 1, Name: "List", Description: new(string)},
		{ID: 2, Name: "End", Quantity: new(int64)},
	}
	*data[0].Description = "<b>Description</b>"
	*data[1].Description = "Placeholder field<script>alert('XSS Test')</script>"
	*data[2].Quantity = 5
	return model.TodoW2Grid{Status: "success", Records: data}, nil
}

func (ts *Todo) DeleteTodoW2Action(idList []int64) (model.W2GridActionResponse, error) {
	return model.W2GridActionResponse{Status: "success"}, nil
}

func (ts *Todo) UpdateTodoW2Action(t []model.TodoPatchDTO) (model.W2GridActionResponse, error) {
	return model.W2GridActionResponse{Status: "success"}, nil
}

func (ts *Todo) CreateTodoW2Form(t model.TodoDTO) (model.W2FormResponse, error) {
	return model.W2FormResponse{Status: "success"}, nil
}

func (ts *Todo) UpdateTodoW2Form(id int64, t model.TodoDTO) (model.W2FormResponse, error) {
	return model.W2FormResponse{Status: "success"}, nil
}
