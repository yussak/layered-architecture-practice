package handler

import (
	"net/http"
	"net/http/httptest"
	"server/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// go test ./handler -v
func TestListTodos(t *testing.T) {
	// Echoインスタンス作成
	e := echo.New()

	// モックDB作成
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %s", err)
	}
	defer dbMock.Close()

	// グローバルDBをモックに差し替え
	db.DB = dbMock

	// モックDBの期待値設定
	mockRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Task 1").
		AddRow(2, "Task 2")
	mock.ExpectQuery("SELECT id, name FROM todos").WillReturnRows(mockRows)

	// リクエストとレスポンス作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// テスト対象関数を実行
	err = ListTodos(c)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"ID":1,"Name":"Task 1"},{"ID":2,"Name":"Task 2"}]`, rec.Body.String())

	// モックが期待どおり実行されたことを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
