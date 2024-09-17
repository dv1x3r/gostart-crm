package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"gostart-crm/internal/app/endpoint"
	"gostart-crm/internal/app/service"
	"gostart-crm/internal/app/storage/sqlitedb"
	"gostart-crm/internal/app/utils"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

type App struct {
	db     *sqlx.DB
	echo   *echo.Echo
	config utils.Config
	logger zerolog.Logger
}

func New() (*App, error) {
	a := &App{
		echo:   echo.New(),
		config: utils.GetConfig(),
		logger: utils.GetLogger(),
	}

	if a.config.DBDriver == "sqlite3" {
		// mattn/go-sqlite3 (cgo)
		a.db = sqlx.MustConnect(a.config.DBDriver, a.config.DBString+"?_journal=WAL&_fk=1&_busy_timeout=10000")
		a.db.SetMaxOpenConns(1)
	} else if a.config.DBDriver == "sqlite" {
		// modernc.org/sqlite (cgo-free)
		// a.db = sqlx.MustConnect(a.config.DBDriver, a.config.DBString+"?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)&_pragma=busy_timeout(10000)")
		// a.db.SetMaxOpenConns(1)
		panic(errors.New("modernc.org/sqlite (cgo-free) is not supported, switch to mattn/go-sqlite3"))
	} else {
		panic(errors.New(a.config.DBDriver + " (db driver) is not supported, switch to mattn/go-sqlite3"))
	}

	if err := goose.SetDialect(a.config.DBDriver); err != nil {
		panic(err)
	}

	if err := goose.Up(a.db.DB, "./migrations"); err != nil {
		panic(err)
	}

	if a.config.Debug {
		a.echo.Debug = true
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if a.config.StaticPath != "" {
		a.echo.Static("/", a.config.StaticPath)
	}

	a.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRequestID:  true,
		LogStatus:     true,
		LogRemoteIP:   true,
		LogLatency:    true,
		LogHost:       true,
		LogMethod:     true,
		LogURI:        true,
		LogError:      true,
		HandleError:   true,
		LogValuesFunc: utils.LogEchoRequestFunc,
	}))

	a.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: utils.LogEchoRecoverFunc,
	}))

	a.echo.Use(middleware.RequestID())
	a.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			ctx := context.WithValue(c.Request().Context(), "RequestID", requestID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	a.echo.Use(middleware.Secure())
	a.echo.Use(middleware.BodyLimit("30M"))

	if a.config.GZIP {
		a.echo.Use(middleware.Gzip())
	}

	if a.config.CORS {
		a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}))
	}

	a.echo.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookiePath:     "/",
		CookieName:     "csrftoken",
		CookieMaxAge:   86400,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	a.echo.Use(session.Middleware(
		sessions.NewCookieStore([]byte(a.config.SessionKey))),
	)

	a.echo.Validator = utils.GetValidator()

	attributeGroupStorage := sqlitedb.NewAttributeGroup(a.db)
	attributeGroupService := service.NewAttributeGroup(attributeGroupStorage)
	attributeSetStorage := sqlitedb.NewAttributeSet(a.db)
	attributeSetService := service.NewAttributeSet(attributeSetStorage)
	attributeValueStorage := sqlitedb.NewAttributeValue(a.db)
	attributeValueService := service.NewAttributeValue(attributeValueStorage)
	attributeEndpoint := endpoint.NewAttribute(attributeGroupService, attributeSetService, attributeValueService)

	brandStorage := sqlitedb.NewBrand(a.db)
	brandService := service.NewBrand(brandStorage)
	brandEndpoint := endpoint.NewBrand(brandService)

	categoryStorage := sqlitedb.NewCategory(a.db)
	categoryService := service.NewCategory(categoryStorage)
	categoryEndpoint := endpoint.NewCategory(categoryService)

	supplierStorage := sqlitedb.NewSupplier(a.db)
	supplierService := service.NewSupplier(supplierStorage)
	supplierEndpoint := endpoint.NewSupplier(supplierService)

	productStorage := sqlitedb.NewProduct(a.db)
	productService := service.NewProduct(productStorage)
	productStatusStorage := sqlitedb.NewProductStatus(a.db)
	productStatusService := service.NewProductStatus(productStatusStorage)
	productEndpoint := endpoint.NewProduct(productService, productStatusService)

	orderStorage := sqlitedb.NewOrder(a.db)
	orderService := service.NewOrder(orderStorage)
	orderStatusStorage := sqlitedb.NewOrderStatus(a.db)
	orderStatusService := service.NewOrderStatus(orderStatusStorage)
	paymentMethodStorage := sqlitedb.NewPaymentMethod(a.db)
	paymentMethodService := service.NewPaymentMethod(paymentMethodStorage)
	orderEndpoint := endpoint.NewOrder(orderService, orderStatusService, paymentMethodService)

	adminEndpoint := endpoint.NewAdmin(
		attributeEndpoint,
		brandEndpoint,
		categoryEndpoint,
		supplierEndpoint,
		productEndpoint,
		orderEndpoint,
	)

	adminEndpoint.Register(a.echo.Group("", demoAuthorizationMiddleware(a.config.ReadOnly)))

	filterStorage := sqlitedb.NewFilter(a.db)
	filterService := service.NewFilter(filterStorage)

	clientEndpoint := endpoint.NewClient(categoryService, productService, filterService)
	clientEndpoint.Register(a.echo.Group("/client", demoAuthorizationMiddleware(a.config.ReadOnly)))

	return a, nil
}

func demoAuthorizationMiddleware(readOnly bool) echo.MiddlewareFunc {
	config := struct{ Skipper middleware.Skipper }{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/login/"
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if readOnly && c.Request().Method != "GET" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Demo works in read-only mode"})
			}

			userSession, _ := session.Get("sessionid", c)
			if userSession.Values["authenticated"] == true {
				return next(c)
			}

			if c.Path() == "/" || c.Path() == "/client" {
				return c.Redirect(http.StatusFound, "/login/")
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Not authorized"})
			}
		}
	}
}

func (a *App) MustRun() {
	var err error

	if strings.HasPrefix(a.config.ServerAddress, "tcp://") {

		address := strings.TrimPrefix(a.config.ServerAddress, "tcp://")
		a.echo.Listener, err = net.Listen("tcp", address)
		if err != nil {
			a.logger.Fatal().Str("address", a.config.ServerAddress).Err(err).Msg("failed to net.Listen")
		}

	} else if strings.HasPrefix(a.config.ServerAddress, "unix://") {

		address := strings.TrimPrefix(a.config.ServerAddress, "unix://")
		if _, err = os.Stat(address); err == nil {
			if err = os.Remove(address); err != nil {
				a.logger.Fatal().Str("address", a.config.ServerAddress).Err(err).Msg("unable to remove existing socket file")
			}
		}

		a.echo.Listener, err = net.Listen("unix", address)
		if err != nil {
			a.logger.Fatal().Str("address", a.config.ServerAddress).Err(err).Msg("failed to net.Listen")
		}

		if err = os.Chmod(address, 0777); err != nil {
			a.logger.Fatal().Str("address", a.config.ServerAddress).Err(err).Msg("unable to chmod .sock permissions")
		}

	} else {
		a.logger.Fatal().Str("address", a.config.ServerAddress).Msg("unsupported address format")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err = a.echo.StartServer(new(http.Server))
		if err != nil && err != http.ErrServerClosed {
			a.logger.Fatal().Err(err).Msg("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echo.Shutdown(ctx); err != nil {
		a.logger.Fatal().Err(err)
	}
}
