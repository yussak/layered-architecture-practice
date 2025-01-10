package handler

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

// TODO:package uiに変えるかも？
func ListTodos(c echo.Context) error {
	todos, err := application.GetTodos()
	if err != nil {
		return c.String(http.StatusInternalServerError, "データ取得エラー")
	}

	return c.JSON(http.StatusOK, todos)
}
