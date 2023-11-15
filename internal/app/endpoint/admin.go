package endpoint

import (
	"encoding/json"
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

func (h *Admin) Register(c *echo.Echo) {
	g := c.Group("/admin")
	g.GET("", h.GetRoot)
	g.GET("/todo", h.GetTodoComponent)
	g.GET("/todo/grid", h.GetTodoGrid)
	g.POST("/todo/grid/delete", h.DeleteTodo)
	g.POST("/todo/grid/patch", h.PatchTodo)
	g.GET("/todo/form", h.GetTodoForm)
	g.POST("/todo/form", h.PostTodoForm)
}

func (h *Admin) GetRoot(ctx echo.Context) error {
	cmp := component.Admin("w2go - admin")
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Admin) GetTodoComponent(ctx echo.Context) error {
	cmp := component.AdminTodo()
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Admin) GetTodoGrid(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	err := json.Unmarshal([]byte(c.QueryParam("request")), req)
	if err != nil {
		return err
	}

	res, err := h.todoService.GetTodoW2Grid(*req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) DeleteTodo(c echo.Context) error {
	req := &model.W2GridDeleteRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.todoService.DeleteTodoW2Action(req.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Admin) PatchTodo(c echo.Context) error {
	req := &model.TodoW2PatchRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.todoService.PatchTodoW2Action(req.Changes)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Admin) GetTodoForm(c echo.Context) error {
	return nil
}

func (h *Admin) PostTodoForm(c echo.Context) error {
	form := &model.TodoW2Form{}
	if err := c.Bind(form); err != nil {
		return err
	}

	res, err := h.todoService.UpsertTodoW2Form(form.RecID, form.Record)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
