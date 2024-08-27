package endpoint

import (
	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Admin struct {
	staticVersion string
}

func NewAdmin(staticVersion string) *Admin {
	return &Admin{
		staticVersion: staticVersion,
	}
}

func (ep *Admin) Register(c *echo.Echo) {
	c.GET("/", func(c echo.Context) error {
		cp := component.CoreParams{
			DebugMode:     utils.GetConfig().Debug,
			CsrfToken:     c.Get("csrf").(string),
			StaticVersion: ep.staticVersion,
			Title:         "Admin | Demo CRM",
		}
		return render(c, component.Admin(cp))
	})
}
