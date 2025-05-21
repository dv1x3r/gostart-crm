package template

import (
	"net/http"
	"strconv"

	"gostart-crm/internal/catalog/category"
	"gostart-crm/internal/catalog/product"
	productstatus "gostart-crm/internal/catalog/product_status"

	adminPage "gostart-crm/internal/template/page/admin"
	catalogPage "gostart-crm/internal/template/page/catalog"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	categoryService      *category.Service
	productService       *product.Service
	productStatusService *productstatus.Service
}

func NewHandler(
	categoryService *category.Service,
	productService *product.Service,
	productStatusService *productstatus.Service,
) *Handler {
	return &Handler{
		categoryService:      categoryService,
		productService:       productService,
		productStatusService: productStatusService,
	}
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/", h.getAdminDashboard)
	g.GET("/client", h.getCatalog)
	g.GET("/client/catalog/*/", h.getCatalog)
}

func (h *Handler) getAdminDashboard(c echo.Context) error {
	return Render(c, http.StatusOK, adminPage.Dashboard(c))
}

func (h *Handler) getCatalog(c echo.Context) error {
	var props catalogPage.ListProps
	var err error

	selectedCategory, err := h.categoryService.GetByURL(c.Request().Context(), c.Param("*"))
	if err != nil {
		return err
	}

	props.MainProps.CategoryName = selectedCategory.Name.V

	if props.MainProps.CategoryTree, err = h.categoryService.FetchTree(c.Request().Context(), true); err != nil {
		return err
	}

	if props.Statuses, err = h.productStatusService.FetchAll(c.Request().Context()); err != nil {
		return err
	}

	props.Request.Limit = 5
	props.Request.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	props.Request.Search = c.QueryParam("search")
	props.Request.SearchBy = c.QueryParam("search-by")
	props.Request.Filters = c.QueryParam("filters")

	if props.Products, err = h.productService.FetchList(c.Request().Context(), props.Request, selectedCategory.ID); err != nil {
		return err
	}

	if props.Request.Offset != 0 {
		return Render(c, http.StatusOK, catalogPage.Rows(c, props))
	}

	return Render(c, http.StatusOK, catalogPage.List(c, props))
}
