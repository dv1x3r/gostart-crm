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
		c.Logger().Warn(err)
		return c.JSON(http.StatusBadRequest, model.W2GridDataResponse[any, any]{Status: "error", Message: "invalid request parameters"})
	}

	res, err := h.todoService.GetTodoW2Grid(*req)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Admin) PostTodoData(c echo.Context) error {
	// if form.Cmd == "save" {
	// } else if form.Cmd == "delete" {
	// }
	// return c.JSON(http.StatusBadRequest, model.W2FormResponse{Status: "error", Message: fmt.Sprintf("Unknown command (%s)", form.Cmd)})
	return nil
}

func (h *Admin) PostTodoForm(c echo.Context) error {
	form := &model.TodoW2Form{}
	if err := c.Bind(form); err != nil {
		c.Logger().Warn(err)
		return c.JSON(http.StatusBadRequest, model.W2FormResponse{Status: "error", Message: "incorrect form mapping"})
	}

	if form.RecID == 0 {
		if res, err := h.todoService.AddTodoW2Form(form.Record); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
	} else {
		if res, err := h.todoService.UpdateTodoW2Form(form.RecID, form.Record); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
	}

	return c.JSON(http.StatusOK, model.W2FormResponse{Status: "success"})
}
