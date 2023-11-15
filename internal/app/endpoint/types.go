package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetTodoW2Grid(model.W2GridDataRequest) (model.TodoW2GridResponse, error)
	DeleteTodoW2Action([]int64) (model.W2BaseResponse, error)
	PatchTodoW2Action([]model.TodoPatchDTO) (model.W2BaseResponse, error)
	UpsertTodoW2Form(int64, model.TodoDTO) (model.W2BaseResponse, error)
}
