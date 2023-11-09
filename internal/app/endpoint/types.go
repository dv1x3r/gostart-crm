package endpoint

import "w2go/internal/app/model"

type TodoService interface {
	GetWGrid() model.WGrid[model.Todo]
}
