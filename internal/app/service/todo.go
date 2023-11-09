package service

import "w2go/internal/app/model"

type Todo struct {
}

func NewTodo() *Todo {
	return &Todo{}
}

func (ts *Todo) GetList() []model.Todo {
	return []model.Todo{
		{ID: 0, Name: "Demo", Description: "Description"},
		{ID: 1, Name: "List", Description: "Placeholder field"},
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
