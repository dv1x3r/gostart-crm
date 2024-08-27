package endpoint

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func renderWithStatus(c echo.Context, component templ.Component, status int) error {
	c.Response().Status = status
	return component.Render(c.Request().Context(), c.Response())
}
