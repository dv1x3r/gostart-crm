package app

import (
	"net/http"

	"w2go/internal/app/endpoint"
	"w2go/internal/app/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	echo          *echo.Echo
	indexEndpoint *endpoint.Index
	adminEndpoint *endpoint.Admin
	todoService   *service.Todo
}

func New() (*App, error) {
	a := &App{}
	a.echo = echo.New()
	a.echo.Debug = true

	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())

	a.echo.Use(middleware.Gzip())
	a.echo.Use(middleware.Secure())
	a.echo.Use(middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			CookiePath:     "/",
			CookieSameSite: http.SameSiteLaxMode,
		},
	))

	a.echo.Static("/assets", "dist")

	a.todoService = &service.Todo{}

	a.indexEndpoint = endpoint.NewIndex(a.todoService)
	a.echo.GET("/", a.indexEndpoint.Get)

	a.adminEndpoint = endpoint.NewAdmin(a.todoService)
	ag := a.echo.Group("/admin")
	ag.GET("", a.adminEndpoint.Get)
	ag.GET("/todo", a.adminEndpoint.GetTodo)
	ag.GET("/todo/data", a.adminEndpoint.GetTodoData)

	a.echo.GET("/users/:id", func(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	a.echo.GET("/show", func(c echo.Context) error {
		// Get team and member from the query string
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, "team:"+team+", member:"+member)
	})

	return a, nil
}

func (a *App) MustRun() {
	a.echo.Logger.Fatal(a.echo.Start("localhost:1323"))
}
