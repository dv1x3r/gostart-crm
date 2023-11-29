package endpoint

import (
	"encoding/json"
	"net/http"

	"w2go/internal/app/component"
	"w2go/internal/app/model"

	"github.com/labstack/echo/v4"
)

type TodoAdminService interface {
	GetTodoW2Grid(model.W2GridDataRequest) (model.TodoW2GridResponse, error)
	DeleteTodoW2Action(model.W2GridDeleteRequest) (model.W2BaseResponse, error)
	PatchTodoW2Action(model.TodoW2PatchRequest) (model.W2BaseResponse, error)
	GetTodoW2Form(model.TodoW2FormRequest) (model.TodoW2FormResponse, error)
	UpsertTodoW2Form(model.TodoW2FormRequest) (model.W2BaseResponse, error)
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
	g.GET("/todo", h.GetTodoComponent)
	g.GET("/todo/grid", h.GetTodoGrid)
	g.POST("/todo/grid/delete", h.DeleteTodoGrid)
	g.POST("/todo/grid/patch", h.PatchTodoGrid)
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.GetTodoW2Grid(*req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Admin) DeleteTodoGrid(c echo.Context) error {
	req := &model.W2GridDeleteRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.DeleteTodoW2Action(*req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Admin) PatchTodoGrid(c echo.Context) error {
	req := &model.TodoW2PatchRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": "bad request",
		})
	}

	res, err := h.todoService.PatchTodoW2Action(*req)
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

	res, err := h.todoService.GetTodoW2Form(*req)
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

	res, err := h.todoService.UpsertTodoW2Form(*req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
