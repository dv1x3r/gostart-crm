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
	g.GET("/todo", h.GetTodo)
	g.GET("/todo/data", h.GetTodoData)
	g.POST("/todo/data", h.PostTodoData)
	g.GET("/todo/form", h.GetTodoForm)
	g.POST("/todo/form", h.PostTodoForm)
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

func (h *Admin) PostTodoData(c echo.Context) error {
	action := &model.TodoW2Action{}
	if err := c.Bind(action); err != nil {
		return err
	}

	if action.Action == "save" {
		res, err := h.todoService.UpdateTodoW2Action(action.Changes)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	} else if action.Action == "delete" {
		res, err := h.todoService.DeleteTodoW2Action(action.ID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusBadRequest, model.W2GridActionResponse{Status: "error", Message: "invalid action"})
}

func (h *Admin) GetTodoForm(c echo.Context) error {
	return nil
}

func (h *Admin) PostTodoForm(c echo.Context) error {
	form := &model.TodoW2Form{}
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.RecID == 0 {
		res, err := h.todoService.CreateTodoW2Form(form.Record)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res, err := h.todoService.UpdateTodoW2Form(form.RecID, form.Record)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}
