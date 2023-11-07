package handlers

import (
	"w2go/components"

	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) Get(c echo.Context) error {
	component := components.Index("w2go - index")
	return component.Render(c.Request().Context(), c.Response().Writer)
}
