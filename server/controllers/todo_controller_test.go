package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// go test ./controllers -v
func TestDeleteTodo(t *testing.T) {
	// テスト用のEchoインスタンス作成
	e := echo.New()

	// sqlmock を使用してデータベースをモック化
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %s", err)
	}
	defer dbMock.Close()

	// モックDBをグローバルな `db.DB` に設定
	db.DB = dbMock

	// テストケース1: 正常な削除リクエスト
	t.Run("正常系: IDが指定されている場合", func(t *testing.T) {
		// モックDBの期待するクエリと結果を定義
		mock.ExpectExec("DELETE FROM todos WHERE id = \\$1").
			WithArgs("123").
			WillReturnResult(sqlmock.NewResult(0, 1)) // 1行削除される

		// リクエスト作成
		req := httptest.NewRequest(http.MethodGet, "/delete?id=123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// テスト対象の関数を実行
		err := DeleteTodo(c)
		assert.NoError(t, err)

		// ステータスコードが期待通り
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		assert.Equal(t, "/", rec.Header().Get("Location"))

		// モックDBの期待通りに動作したかを検証
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// テストケース2: IDが空の場合
	t.Run("異常系: IDが空", func(t *testing.T) {
		// リクエスト作成
		req := httptest.NewRequest(http.MethodGet, "/delete?id=", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// テスト対象の関数を実行
		err := DeleteTodo(c)
		assert.NoError(t, err)

		// ステータスコードとエラーメッセージが期待通り
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "IDが空です", rec.Body.String())
	})

	// テストケース3: データベースエラーの場合
	t.Run("異常系: データベースエラー", func(t *testing.T) {
		// モックDBの期待するクエリと結果を定義
		mock.ExpectExec("DELETE FROM todos WHERE id = \\$1").
			WithArgs("123").
			WillReturnError(fmt.Errorf("データベースエラー"))

		// リクエスト作成
		req := httptest.NewRequest(http.MethodGet, "/delete?id=123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// テスト対象の関数を実行
		err := DeleteTodo(c)
		assert.NoError(t, err)

		// ステータスコードとエラーメッセージが期待通り
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "データベースエラー", rec.Body.String())

		// モックDBの期待通りに動作したかを検証
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
