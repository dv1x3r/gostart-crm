package template

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, statusCode int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	err := component.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	return c.HTML(statusCode, buf.String())
}
