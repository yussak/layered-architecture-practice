package routes

import (
	"server/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return handler.ListTodosHandler(c)
	})
	e.POST("/add", func(c echo.Context) error {
		return handler.AddTodoHander(c)
	})
	e.DELETE("/delete", func(c echo.Context) error {
		return handler.DeleteTodoHandler(c)
	})
}
