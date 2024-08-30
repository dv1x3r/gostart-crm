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
	supplierEndpoint  *Supplier
	orderEndpoint     *Order
}

func NewAdmin(
	staticVersion string,
	attributeEndpoint *Attribute,
	brandEndpoint *Brand,
	supplierEndpoint *Supplier,
	orderEndpoint *Order,

) *Admin {
	return &Admin{
		staticVersion:     staticVersion,
		attributeEndpoint: attributeEndpoint,
		brandEndpoint:     brandEndpoint,
		supplierEndpoint:  supplierEndpoint,
		orderEndpoint:     orderEndpoint,
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

	ep.attributeEndpoint.Register(c.Group("/attribute"))
	ep.brandEndpoint.Register(c.Group("/brand"))
	ep.supplierEndpoint.Register(c.Group("/supplier"))
	ep.orderEndpoint.Register(c.Group("/order"))
}
