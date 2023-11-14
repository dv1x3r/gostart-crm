package service

import "w2go/internal/app/model"

type Todo struct {
}

func (ts *Todo) GetTodoW2Grid(req model.W2GridDataRequest) (model.TodoW2Grid, error) {
	data := []model.TodoDTO{
		{ID: 0, Name: "Demo", Description: "<b>Description</b>"},
		{ID: 1, Name: "List", Description: "Placeholder field<script>alert('XSS Test')</script>"},
		{ID: 2, Name: "End", Description: "the end of the list"},
	}
	return model.TodoW2Grid{Status: "success", Records: data}, nil
}

func (ts *Todo) AddTodoW2Form(t model.TodoDTO) (model.W2FormResponse, error) {
	return model.W2FormResponse{Status: "success"}, nil
}

func (ts *Todo) UpdateTodoW2Form(id int64, t model.TodoDTO) (model.W2FormResponse, error) {
	return model.W2FormResponse{Status: "success"}, nil
}
