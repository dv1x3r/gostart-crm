package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetTodoW2Grid(model.W2GridDataRequest) (model.TodoW2GridResponse, error)
	DeleteTodoW2Action(model.W2GridDeleteRequest) (model.W2BaseResponse, error)
	PatchTodoW2Action(model.TodoW2PatchRequest) (model.W2BaseResponse, error)
	GetTodoW2Form(model.TodoW2FormRequest) (model.TodoW2FormResponse, error)
	UpsertTodoW2Form(model.TodoW2FormRequest) (model.W2BaseResponse, error)
}
