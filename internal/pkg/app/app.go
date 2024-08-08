package app

import (
	"net/http"
	"os"

	"gocraft-crm/internal/app/endpoint"
	"gocraft-crm/internal/app/service"
	"gocraft-crm/internal/app/storage/sqlitedb"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	echo          *echo.Echo
	db            *sqlx.DB
	indexEndpoint *endpoint.Index
	adminEndpoint *endpoint.Admin
	todoStorage   *sqlitedb.Todo
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

	dbDriver, dbString := os.Getenv("GOOSE_DRIVER"), os.Getenv("GOOSE_DBSTRING")
	if dbDriver == "sqlite3" {
		dbString += "?_journal=WAL&_fk=1"
	}
	a.db = sqlx.MustConnect(dbDriver, dbString)

	a.todoStorage = sqlitedb.NewTodo(a.db)
	a.todoService = service.NewTodo(a.todoStorage)

	a.indexEndpoint = endpoint.NewIndex(a.todoService)
	a.echo.GET("/", a.indexEndpoint.GetRoot)

	a.adminEndpoint = endpoint.NewAdmin(a.todoService)
	a.adminEndpoint.Register(a.echo)

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
