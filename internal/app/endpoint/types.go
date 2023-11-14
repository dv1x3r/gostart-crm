package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetTodoW2Grid(model.W2GridDataRequest) (model.TodoW2Grid, error)
	DeleteTodoW2Action([]int64) (model.W2GridActionResponse, error)
	UpdateTodoW2Action([]model.TodoPatchDTO) (model.W2GridActionResponse, error)
	CreateTodoW2Form(model.TodoDTO) (model.W2FormResponse, error)
	UpdateTodoW2Form(int64, model.TodoDTO) (model.W2FormResponse, error)
}
