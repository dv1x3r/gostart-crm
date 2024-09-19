package endpoint

import (
	"strconv"

	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/service"
	"gostart-crm/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Client struct {
	categoryService *service.Category
	productService  *service.Product
	filterService   *service.Filter
}

func NewClient(
	categoryService *service.Category,
	productService *service.Product,
	filterService *service.Filter,
) *Client {
	return &Client{
		categoryService: categoryService,
		productService:  productService,
		filterService:   filterService,
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
		DebugMode: utils.GetConfig().Debug,
		CsrfToken: c.Get("csrf").(string),
		Title:     "Client | Demo CRM",
		GoogleTag: utils.GetConfig().GoogleTag,
	}

	if params.CategoryTree, err = ep.categoryService.FetchTree(ctx); err != nil {
		return err
	}

	params.SelectedCategoryID, _ = strconv.ParseInt(c.QueryParam("category"), 10, 64)
	appliedFilters := ep.filterService.ParseAppliedFilters(c.QueryParam("filters"))
	if params.SelectedCategoryID != 0 {
		params.FilterFacets, err = ep.filterService.GetFacetsByCategoryID(ctx, params.SelectedCategoryID, appliedFilters)
		if err != nil {
			return err
		}
	}

	params.Products, err = ep.productService.FetchAvailableByCategoryID(ctx, params.SelectedCategoryID, appliedFilters)
	if err != nil {
		return err
	}

	return render(c, component.ClientMainPage(params))
}
