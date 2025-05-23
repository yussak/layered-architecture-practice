package ui

// ui層の役割
// リクエストを受け取り、形式的なチェックを行う
// アプリの仕様に基づくチェックはここではなくapplication層で行う

// 以前req.Name == "" のチェックをuiで書いていたが、Nameが空でも技術的には問題がない。でもアプリ側で問題がある。なのでapplicationで書くのか

import (
	"net/http"
	"server/application"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Service application.TodoService
}

// TodoHandlerのServiceを通じてapplicationのGetTodos()にアクセスするように変更し、直接application.GetTodos()にアクセスしなくなった
// それによってモックしやすくなる
func (h *TodoHandler) HandleGetTodos(c echo.Context) error {
	todos, err := h.Service.GetTodos()
	if err != nil {
		return c.String(http.StatusInternalServerError, "データ取得エラー")
	}

	return c.JSON(http.StatusOK, todos)
}

// リクエスト用のDTOを定義
type AddTodoRequest struct {
	Name string `json:"name"`
}

func HandleAddTodo(c echo.Context) error {
	var req AddTodoRequest

	// JSONボディをバインド
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの形式が正しくありません")
	}

	todo, err := application.CreateTodo(req.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	return c.JSON(http.StatusOK, todo)
}

func HandleDeleteTodo(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "IDが空です")
	}

	// データベースから削除
	err := application.DeleteTodo(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベースエラー")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
