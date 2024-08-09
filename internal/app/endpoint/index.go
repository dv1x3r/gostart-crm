package endpoint

import (
	"context"

	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/model"

	"github.com/labstack/echo/v4"
)

type TodoService interface {
	GetTodoList(context.Context) ([]model.TodoFromDB, error)
}

type Index struct {
	todoService TodoService
}

func NewIndex(ts TodoService) *Index {
	return &Index{
		todoService: ts,
	}
}

func (h *Index) GetRoot(c echo.Context) error {
	todos, err := h.todoService.GetTodoList(c.Request().Context())
	if err != nil {
		return err
	} else {
		cmp := component.Index("w2go - index", todos)
		return cmp.Render(c.Request().Context(), c.Response().Writer)
	}
}
