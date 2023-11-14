package endpoint

import (
	"net/http"

	"w2go/internal/app/component"
	"w2go/internal/app/model"

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
	grid, _ := h.todoService.GetList()
	return c.JSON(http.StatusOK, grid)
}

func (h *Admin) PostTodoData(c echo.Context) error {

	// form := &model.W2FormSubmit[model.Todo, int64]{}
	// if err := c.Bind(form); err != nil {
	// 	return c.JSON(http.StatusBadRequest, model.W2FormResponse{Status: "error", Message: "incorrect data"})
	// }
	// c.Logger().Debug(fmt.Printf("%+v\n", form))
	_ = h.todoService.Create(model.Todo{})
	return c.JSON(http.StatusOK, "")
}
