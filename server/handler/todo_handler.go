package handler

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

func ListTodosHandler(c echo.Context) error {
	todos, err := application.GetTodos()
	if err != nil {
		return c.String(http.StatusInternalServerError, "データ取得エラー")
	}

	return c.JSON(http.StatusOK, todos)
}

// リクエスト用のDTOを定義
type AddTodoRequest struct {
	Name string `json:"name"`
}

// TODO:AddTodoHanderにしたい
func AddTodo(c echo.Context) error {
	var req AddTodoRequest

	// JSONボディをバインド
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}
	if req.Name == "" {
		return c.String(http.StatusBadRequest, "TODO名が空です")
	}

	newTodo, err := application.CreateTodo(req.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "a")
	}

	return c.JSON(http.StatusOK, newTodo)
}
