package endpoint

import (
	"net/http"

	"w2go/internal/app/component"

	"github.com/labstack/echo/v4"
)

type Admin struct {
	todoService TodoService
}

func NewAdmin(s TodoService) *Admin {
	return &Admin{
		todoService: s,
	}
}

func (h *Admin) GetRoot(ctx echo.Context) error {
	cmp := component.Admin("w2go - admin")
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Admin) GetTodo(ctx echo.Context) error {
	cmp := component.AdminTodo()
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Admin) GetTodoData(c echo.Context) error {
	grid := h.todoService.GetWGrid()
	return c.JSON(http.StatusOK, grid)
}
