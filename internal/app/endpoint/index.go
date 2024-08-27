package endpoint

import (
	"gostart-crm/internal/app/component"

	"github.com/labstack/echo/v4"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (h *Index) GetRoot(ctx echo.Context) error {
	cp := component.CoreParams{
		Title: "w2go - admin",
	}
	cmp := component.Client(cp)
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}
