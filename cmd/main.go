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
		component := views.Index("w2go - index")
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	admin := e.Group("/admin")

	admin.GET("", func(c echo.Context) error {
		component := views.Admin("w2go - admin")
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	admin.GET("/todo", func(c echo.Context) error {
		component := views.AdminTodo()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	admin.GET("/todo/data", getAdminTodo)

	e.Logger.Fatal(e.Start("localhost:1323"))
}

func getAdminTodo(c echo.Context) error {
	var todos = []Todo{}

	for i := 0; i < 100; i++ {
		todos = append(todos, Todo{ID: int64(i), Name: "todo name", Description: "todo description"})
	}

	return c.JSON(http.StatusOK, todos)
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
