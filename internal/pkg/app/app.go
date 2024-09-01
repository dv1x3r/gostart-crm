package app

import (
	"context"
	"fmt"
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

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	_ "github.com/mattn/go-sqlite3"
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

	a.echo.Validator = utils.GetValidator()

	if a.config.DBDriver == "sqlite3" {
		a.db = sqlx.MustConnect(a.config.DBDriver, a.config.DBString+"?_journal=WAL&_fk=1&_busy_timeout=10000")
		a.db.SetMaxOpenConns(1)
	} else if a.config.DBDriver == "sqlite" {
		a.db = sqlx.MustConnect(a.config.DBDriver, a.config.DBString+"?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)&_pragma=busy_timeout(10000)")
		a.db.SetMaxOpenConns(1)
	} else {
		a.db = sqlx.MustConnect(a.config.DBDriver, a.config.DBString)
	}

	if a.config.Debug {
		a.echo.Debug = true
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.TimeFieldFormat = time.TimeOnly
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = time.RFC3339
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

	staticVersion := fmt.Sprint(time.Now().Unix())

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

	orderStorage := sqlitedb.NewOrder(a.db)
	orderService := service.NewOrder(orderStorage)

	orderStatusStorage := sqlitedb.NewOrderStatus(a.db)
	orderStatusService := service.NewOrderStatus(orderStatusStorage)

	paymentMethodStorage := sqlitedb.NewPaymentMethod(a.db)
	paymentMethodService := service.NewPaymentMethod(paymentMethodStorage)

	orderEndpoint := endpoint.NewOrder(orderService, orderStatusService, paymentMethodService)

	adminEndpoint := endpoint.NewAdmin(
		staticVersion,
		attributeEndpoint,
		brandEndpoint,
		categoryEndpoint,
		supplierEndpoint,
		orderEndpoint,
	)

	adminEndpoint.Register(a.echo)

	return a, nil
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
