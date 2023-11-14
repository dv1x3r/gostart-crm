package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetList() ([]model.Todo, error)
	Create(model.Todo) error
}
