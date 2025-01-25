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

// TodoHandlerのServiceを通じてapplicationのGetTodos()にアクセスするように変更し、直接application.GetTodos()にアクセスしなくなった
// それによってモックしやすくなる
// 現状ui->appのように一つだけに依存しているため引数で受け取っているが、複数の依存がある場合は構造体にしてその中に入れるほうが見やすいのでそうする予定
func HandleGetTodos(todoService application.TodoService, c echo.Context) error {
	todos, err := todoService.GetTodos()
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
