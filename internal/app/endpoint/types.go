package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetTodoW2Grid(model.W2GridDataRequest) (model.TodoW2Grid, error)
	AddTodoW2Form(model.TodoDTO) (model.W2FormResponse, error)
	UpdateTodoW2Form(int64, model.TodoDTO) (model.W2FormResponse, error)
}
