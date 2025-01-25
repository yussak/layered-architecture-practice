package routes

import (
	"server/application"
	"server/ui"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, todoService application.TodoService) {
	e.GET("/", func(c echo.Context) error {
		return ui.HandleGetTodos(todoService, c)
	})
	e.POST("/add", func(c echo.Context) error {
		return ui.HandleAddTodo(c)
	})
	e.DELETE("/delete", func(c echo.Context) error {
		return ui.HandleDeleteTodo(c)
	})
}
