package routes

import (
	"server/ui"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return ui.HandleGetTodos(c)
	})
	e.POST("/add", func(c echo.Context) error {
		return ui.HandleAddTodo(c)
	})
	e.DELETE("/delete", func(c echo.Context) error {
		return ui.HandleDeleteTodo(c)
	})
}
