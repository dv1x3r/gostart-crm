package endpoint

import (
	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Admin struct {
	staticVersion     string
	attributeEndpoint *Attribute
	brandEndpoint     *Brand
	categoryEndpoint  *Category
	supplierEndpoint  *Supplier
	productEndpoint   *Product
	orderEndpoint     *Order
}

func NewAdmin(
	staticVersion string,
	attributeEndpoint *Attribute,
	brandEndpoint *Brand,
	categoryEndpoint *Category,
	supplierEndpoint *Supplier,
	productEndpoint *Product,
	orderEndpoint *Order,

) *Admin {
	return &Admin{
		staticVersion:     staticVersion,
		attributeEndpoint: attributeEndpoint,
		brandEndpoint:     brandEndpoint,
		categoryEndpoint:  categoryEndpoint,
		supplierEndpoint:  supplierEndpoint,
		productEndpoint:   productEndpoint,
		orderEndpoint:     orderEndpoint,
	}
}

func (ep *Admin) Register(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		cp := component.CoreParams{
			DebugMode:     utils.GetConfig().Debug,
			CsrfToken:     c.Get("csrf").(string),
			StaticVersion: ep.staticVersion,
			Title:         "Admin | Demo CRM",
		}
		return render(c, component.Admin(cp))
	})

	ep.attributeEndpoint.Register(g.Group("/attribute"))
	ep.brandEndpoint.Register(g.Group("/brand"))
	ep.categoryEndpoint.Register(g.Group("/category"))
	ep.supplierEndpoint.Register(g.Group("/supplier"))
	ep.productEndpoint.Register(g.Group("/product"))
	ep.orderEndpoint.Register(g.Group("/order"))
}
