package service

import "w2go/internal/app/model"

type Todo struct {
}

func (ts *Todo) Create(t model.Todo) error {
	return nil
}

func (ts *Todo) GetList() ([]model.Todo, error) {
	return []model.Todo{
		{ID: 0, Name: "Demo", Description: "<b>Description</b>"},
		{ID: 1, Name: "List", Description: "Placeholder field<script>alert('XSS Test')</script>"},
		{ID: 2, Name: "End", Description: "the end of the list"},
	}, nil
}

// func (ts *Todo) GetW2Grid() model.W2GridResponseDTO[model.Todo, any] {
// 	data := ts.GetList()
// 	return model.W2GridResponseDTO[model.Todo, any]{
// 		Status:  "success",
// 		Records: data,
// 	}
// }

func (ts *Todo) NewTodo() (model.Todo, error) {
	return model.Todo{
		// Status: "success",
	}, nil
}
