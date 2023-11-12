package service

import "w2go/internal/app/model"

type Todo struct {
}

func NewTodo() *Todo {
	return &Todo{}
}

func (ts *Todo) GetList() []model.Todo {
	return []model.Todo{
		{ID: 0, Name: "Demo", Description: "<b>Description</b>"},
		{ID: 1, Name: "List", Description: "Placeholder field<script>alert('XSS Test')</script>"},
		{ID: 2, Name: "End", Description: "the end of the list"},
	}
}

func (ts *Todo) GetW2Grid() model.W2Grid[model.Todo] {
	data := ts.GetList()
	return model.W2Grid[model.Todo]{
		Status:  "success",
		Records: data,
	}
}

func (ts *Todo) NewTodo() model.W2FormResponse {
	return model.W2FormResponse{
		Status: "success",
	}
}
