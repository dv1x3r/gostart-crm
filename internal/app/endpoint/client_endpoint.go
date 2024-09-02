package endpoint

import (
	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Client struct {
	staticVersion string
}

func NewClient(
	staticVersion string,

) *Client {
	return &Client{
		staticVersion: staticVersion,
	}
}

func (ep *Client) Register(g *echo.Group) {
	g.GET("", ep.GetRoot)
}

func (ep *Client) GetRoot(c echo.Context) error {
	// ctx := c.Request().Context()
	// var params component.MainPageParams
	// var err error

	cp := component.CoreParams{
		DebugMode:     utils.GetConfig().Debug,
		CsrfToken:     c.Get("csrf").(string),
		StaticVersion: ep.staticVersion,
		Title:         "Client | Demo CRM",
	}

	return render(c, component.Client(cp))
}
