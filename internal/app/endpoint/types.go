package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetW2Grid() model.W2Grid[model.Todo]
	NewTodo() model.W2FormResponse
}
