package main

import (
	"net/http"
	"w2go/pkg/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "dist")

	e.GET("/", func(c echo.Context) error {
		component := views.Index("root")
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.GET("/todos", func(c echo.Context) error {
		component := views.Todos()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/api/todo", func(c echo.Context) error {
		var todos = []Todo{
			{ID: 1, Name: "Test", Description: "Some text"},
			{ID: 1, Name: "Test", Description: "Some text"},
		}

		for i := 0; i < 10000; i++ {
			todos = append(todos, Todo{ID: int64(i), Name: "New", Description: "desc"})

		}

		return c.JSON(http.StatusOK, todos)
	})

	e.Logger.Fatal(e.Start("localhost:1323"))
}

type Todo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
