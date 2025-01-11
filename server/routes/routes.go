package routes

import (
	"server/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return handler.HandleGetTodos(c)
	})
	e.POST("/add", func(c echo.Context) error {
		return handler.HandleAddTodo(c)
	})
	e.DELETE("/delete", func(c echo.Context) error {
		return handler.HandleDeleteTodo(c)
	})
}
