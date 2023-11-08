package endpoint

import (
	"w2go/internal/app/component"

	"github.com/labstack/echo/v4"
)

type Index struct {
	todoService TodoService
}

func NewIndex(ts TodoService) *Index {
	return &Index{
		todoService: ts,
	}
}

func (h *Index) Get(ctx echo.Context) error {
	cmp := component.Index("w2go - index")
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}
