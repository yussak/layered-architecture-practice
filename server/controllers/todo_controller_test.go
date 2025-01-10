package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"server/db"
	"server/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// go test ./controllers -v
// func TestListTodos(t *testing.T) {
// 	// Echoインスタンス作成
// 	e := echo.New()

// 	// モックDB作成
// 	dbMock, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("failed to create sqlmock: %s", err)
// 	}
// 	defer dbMock.Close()

// 	// グローバルDBをモックに差し替え
// 	db.DB = dbMock

// 	// モックDBの期待値設定
// 	mockRows := sqlmock.NewRows([]string{"id", "name"}).
// 		AddRow(1, "Task 1").
// 		AddRow(2, "Task 2")
// 	mock.ExpectQuery("SELECT id, name FROM todos").WillReturnRows(mockRows)

// 	// リクエストとレスポンス作成
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// テスト対象関数を実行
// 	err = ListTodos(c)

// 	// アサーション
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.JSONEq(t, `[{"ID":1,"Name":"Task 1"},{"ID":2,"Name":"Task 2"}]`, rec.Body.String())

// 	// モックが期待どおり実行されたことを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestAddTodoWithSQLMock(t *testing.T) {
	// モックデータベースとモックインターフェースを作成
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// モックデータベースをグローバル変数に設定（テスト用）
	db.DB = mockDB

	// モックデータ
	mockTodo := models.Todo{Name: "Test Todo"}
	mockResponse := models.Todo{ID: 1, Name: "Test Todo"}

	// JSONリクエストボディ作成
	requestBody, _ := json.Marshal(mockTodo)

	// Echoのリクエストとレスポンスを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// INSERTクエリのモック設定
	mock.ExpectQuery("INSERT INTO Todos").
		WithArgs(mockTodo.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// テスト実行
	if assert.NoError(t, AddTodo(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディの検証
		var responseTodo models.Todo
		err := json.Unmarshal(rec.Body.Bytes(), &responseTodo)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, responseTodo)
	}

	// モックの期待値の検証
	assert.NoError(t, mock.ExpectationsWereMet())
}
