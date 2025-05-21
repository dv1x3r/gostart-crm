package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"gostart-crm/internal/config"
	"gostart-crm/internal/lib/db"
	"gostart-crm/internal/lib/logger"
	"gostart-crm/internal/lib/validator"
	"gostart-crm/internal/random/discount"
	"gostart-crm/internal/template"
	"gostart-crm/migrations"
	"gostart-crm/web"

	"gostart-crm/internal/catalog/attribute"
	"gostart-crm/internal/catalog/brand"
	"gostart-crm/internal/catalog/category"
	"gostart-crm/internal/catalog/product"
	"gostart-crm/internal/catalog/product_status"
	"gostart-crm/internal/catalog/supplier"

	"gostart-crm/internal/checkout/order"
	"gostart-crm/internal/checkout/order_status"
	"gostart-crm/internal/checkout/payment_method"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()

	if config.Debug {
		logger.SetLevel(slog.LevelDebug)
	}

	store, err := db.NewStore(config.DBDriver, config.DBString)
	if err != nil {
		logger.GetLogger().Error("unable to connect to the database", "err", err)
		os.Exit(1)
	}

	if err := db.MigrateUp(logger.GetLogger(), config.DBDriver, store.DB(), migrations.MigrationsFS, "."); err != nil {
		logger.GetLogger().Error("unable to migrate the database", "err", err)
		os.Exit(1)
	}

	e := echo.New()
	e.HideBanner = true
	e.Debug = config.Debug
	e.Validator = validator.GetValidator()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogMethod:   true,
		LogHost:     true,
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.GetLogger().Info("REQUEST",
					slog.String("remote_ip", v.RemoteIP),
					slog.String("method", v.Method),
					slog.String("host", v.Host),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("latency", fmt.Sprintf("%s", v.Latency)),
				)
			} else {
				logger.GetLogger().Error("REQUEST_ERROR",
					slog.String("remote_ip", v.RemoteIP),
					slog.String("method", v.Method),
					slog.String("host", v.Host),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("latency", fmt.Sprintf("%s", v.Latency)),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookieName:     "csrftoken",
		CookiePath:     "/",
		CookieMaxAge:   86400,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	e.Use(middleware.BodyLimit("30M"))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	static := e.Group("", addCacheControlHeader("public, max-age=86400"))
	static.StaticFS("/admin.js", echo.MustSubFS(web.WebFS, "dist/admin.js"))
	static.StaticFS("/admin.css", echo.MustSubFS(web.WebFS, "dist/admin.css"))
	static.StaticFS("/client.js", echo.MustSubFS(web.WebFS, "dist/client.js"))
	static.StaticFS("/client.css", echo.MustSubFS(web.WebFS, "dist/client.css"))
	static.StaticFS("/site.webmanifest", echo.MustSubFS(web.WebFS, "static/site.webmanifest"))

	api := e.Group("", readOnlyMiddleware(config.ReadOnly))

	attributeRepo := attribute.NewSQLiteRepository(store)
	attributeService := attribute.NewService(attributeRepo)
	attributeHandler := attribute.NewHandler(attributeService)
	attributeHandler.Register(api.Group("/attribute"))

	brandRepo := brand.NewSQLiteRepository(store)
	brandService := brand.NewService(brandRepo)
	brandHandler := brand.NewHandler(brandService)
	brandHandler.Register(api.Group("/brand"))

	categoryRepo := category.NewSQLiteRepository(store)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)
	categoryHandler.Register(api.Group("/category"))

	supplierRepo := supplier.NewSQLiteRepository(store)
	supplierService := supplier.NewService(supplierRepo)
	supplierHandler := supplier.NewHandler(supplierService)
	supplierHandler.Register(api.Group("/supplier"))

	productRepo := product.NewSQLiteRepository(store)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)
	productHandler.Register(api.Group("/product"))

	productStatusRepo := productstatus.NewSQLiteRepository(store)
	productStatusService := productstatus.NewService(productStatusRepo)
	productStatusHandler := productstatus.NewHandler(productStatusService)
	productStatusHandler.Register(api.Group("/product/status"))

	orderRepo := order.NewSQLiteRepository(store)
	orderService := order.NewService(orderRepo)
	orderHandler := order.NewHandler(orderService)
	orderHandler.Register(api.Group("/order"))

	orderStatusRepo := orderstatus.NewSQLiteRepository(store)
	orderStatusService := orderstatus.NewService(orderStatusRepo)
	orderStatusHandler := orderstatus.NewHandler(orderStatusService)
	orderStatusHandler.Register(api.Group("/order/status"))

	paymentMethodRepo := paymentmethod.NewSQLiteRepository(store)
	paymentMethodService := paymentmethod.NewService(paymentMethodRepo)
	paymentMethodHandler := paymentmethod.NewHandler(paymentMethodService)
	paymentMethodHandler.Register(api.Group("/payment/method"))

	discountHandler := discount.NewHandler()
	discountHandler.Register(e.Group(""))

	templateHandler := template.NewHandler(
		categoryService,
		productService,
		productStatusService,
	)
	templateHandler.Register(e.Group(""))

	go startServer(e, config.ServerAddress)
	select {}
}

func addCacheControlHeader(value string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderCacheControl, value)
			return next(c)
		}
	}
}

func readOnlyMiddleware(readOnly bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if readOnly && c.Request().Method != "GET" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Demo works in read-only mode"})
			}

			return next(c)
		}
	}
}

func startServer(e *echo.Echo, address string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		address = strings.TrimPrefix(address, "tcp://")
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
