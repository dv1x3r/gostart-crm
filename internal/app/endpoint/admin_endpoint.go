package endpoint

import (
	"crypto/subtle"
	"net/http"

	"gostart-crm/internal/app/component"
	"gostart-crm/internal/app/utils"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Admin struct {
	attributeEndpoint *Attribute
	brandEndpoint     *Brand
	categoryEndpoint  *Category
	supplierEndpoint  *Supplier
	productEndpoint   *Product
	orderEndpoint     *Order
}

func NewAdmin(
	attributeEndpoint *Attribute,
	brandEndpoint *Brand,
	categoryEndpoint *Category,
	supplierEndpoint *Supplier,
	productEndpoint *Product,
	orderEndpoint *Order,

) *Admin {
	return &Admin{
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
			DebugMode: utils.GetConfig().Debug,
			CsrfToken: c.Get("csrf").(string),
			Title:     "Admin | Demo CRM",
		}
		return render(c, component.Admin(cp))
	})

	g.GET("/login/", func(c echo.Context) error {
		userSession, _ := session.Get("sessionid", c)
		if userSession.Values["authenticated"] == true {
			// do not show login page for authenticated user
			return c.Redirect(http.StatusFound, "/")
		}

		cp := component.CoreParams{
			DebugMode: utils.GetConfig().Debug,
			CsrfToken: c.Get("csrf").(string),
			Title:     "Login | Demo CRM",
		}
		return render(c, component.AdminLogin(cp))
	})

	g.POST("/login/", func(c echo.Context) error {
		inputEmail := c.FormValue("email")
		inputPassword := c.FormValue("password")

		// use constant time string comparison to prevent timing bruteforce
		if subtle.ConstantTimeCompare([]byte(inputEmail), []byte(utils.GetConfig().AdminLogin)) == 1 &&
			subtle.ConstantTimeCompare([]byte(inputPassword), []byte(utils.GetConfig().AdminPassword)) == 1 {
			c.Response().Header().Add("HX-Refresh", "true")

			userSession, _ := session.Get("sessionid", c)
			userSession.Values["authenticated"] = true
			userSession.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   0, // session
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}

			if err := userSession.Save(c.Request(), c.Response()); err != nil {
				return err
			}

			return c.String(200, "")
		}

		return c.String(401, "Unauthorized")
	}, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(1))))

	g.GET("/logout/", func(c echo.Context) error {
		userSession, _ := session.Get("sessionid", c)
		delete(userSession.Values, "authenticated")
		if err := userSession.Save(c.Request(), c.Response()); err != nil {
			return err
		}
		return c.Redirect(http.StatusFound, "/login/")
	})

	ep.attributeEndpoint.Register(g.Group("/attribute"))
	ep.brandEndpoint.Register(g.Group("/brand"))
	ep.categoryEndpoint.Register(g.Group("/category"))
	ep.supplierEndpoint.Register(g.Group("/supplier"))
	ep.productEndpoint.Register(g.Group("/product"))
	ep.orderEndpoint.Register(g.Group("/order"))
}
