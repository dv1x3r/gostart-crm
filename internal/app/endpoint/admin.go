package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/model"

	"github.com/labstack/echo/v4"
)

type TodoAdminService interface {
	GetTodoW2Grid(context.Context, model.W2GridDataRequest) (model.TodoW2GridResponse, error)
	DeleteTodoW2Grid(context.Context, model.W2GridRemoveRequest) (model.W2BaseResponse, error)
	PatchTodoW2Action(context.Context, model.TodoW2SaveRequest) (model.W2BaseResponse, error)
	GetTodoW2Form(context.Context, model.TodoW2FormRequest) (model.TodoW2FormResponse, error)
	UpsertTodoW2Form(context.Context, model.TodoW2FormRequest) (model.W2BaseResponse, error)
}

type Admin struct {
	todoService TodoAdminService
}

func NewAdmin(s TodoAdminService) *Admin {
	return &Admin{
		todoService: s,
	}
}

func (h *Admin) Register(c *echo.Echo) {
	g := c.Group("/admin")
	g.GET("", h.GetRoot)
	g.GET("/todo/grid", h.GetTodoGrid)
	g.POST("/todo/grid/delete", h.DeleteTodoGrid)
	g.POST("/todo/grid/patch", h.PatchTodoGrid)
	g.GET("/todo/form", h.GetTodoForm)
	g.POST("/todo/form", h.PostTodoForm)
}

func (h *Admin) GetRoot(ctx echo.Context) error {
	cp := component.CoreParams{
		Title: "w2go - admin",
	}
	cmp := component.Admin(cp)
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Admin) GetTodoGrid(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	err := json.Unmarshal([]byte(c.QueryParam("request")), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.GetTodoW2Grid(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) DeleteTodoGrid(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.DeleteTodoW2Grid(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) PatchTodoGrid(c echo.Context) error {
	req := &model.TodoW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.PatchTodoW2Action(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) GetTodoForm(c echo.Context) error {
	req := &model.TodoW2FormRequest{}
	err := json.Unmarshal([]byte(c.QueryParam("request")), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.GetTodoW2Form(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) PostTodoForm(c echo.Context) error {
	req := &model.TodoW2FormRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.UpsertTodoW2Form(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
