package endpoint

import (
	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/service"
	"gostart-crm/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Client struct {
	staticVersion   string
	categoryService *service.Category
}

func NewClient(
	staticVersion string,
	categoryService *service.Category,
) *Client {
	return &Client{
		staticVersion:   staticVersion,
		categoryService: categoryService,
	}
}

func (ep *Client) Register(g *echo.Group) {
	g.GET("", ep.GetRoot)
}

func (ep *Client) GetRoot(c echo.Context) error {
	ctx := c.Request().Context()
	var params component.ClientMainPageParams
	var err error

	params.Core = component.CoreParams{
		DebugMode:     utils.GetConfig().Debug,
		CsrfToken:     c.Get("csrf").(string),
		StaticVersion: ep.staticVersion,
		Title:         "Client | Demo CRM",
	}

	if params.CategoryTree, err = ep.categoryService.FetchTree(ctx); err != nil {
		return err
	}

	return render(c, component.ClientMainPage(params))
}
