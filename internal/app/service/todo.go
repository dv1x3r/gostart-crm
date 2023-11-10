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

func (ts *Todo) GetWGrid() model.WGrid[model.Todo] {
	data := ts.GetList()
	return model.WGrid[model.Todo]{
		Status:  "success",
		Message: "error message",
		// Total:   256,
		Records: data,
	}
}
