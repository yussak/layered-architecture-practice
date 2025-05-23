package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/db"
	"server/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TODO: package事にテスト書きなおす！！

// go test ./ui -v

// モックの定義
type MockTodoService struct{}

func (m *MockTodoService) GetTodos() ([]domain.Todo, error) {
	return []domain.Todo{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
	}, nil
}

func TestHandleGetTodos(t *testing.T) {
	// Echoインスタンス作成
	e := echo.New()

	// リクエストとレスポンス作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// モックサービスを注入
	mockService := &MockTodoService{}
	handler := TodoHandler{Service: mockService}

	// // テスト対象関数を実行
	err := handler.HandleGetTodos(c)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"ID":1,"Name":"Task 1"},{"ID":2,"Name":"Task 2"}]`, rec.Body.String())
}

func TestAddTodoWithSQLMock(t *testing.T) {
	// モックデータベースとモックインターフェースを作成
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// モックデータベースをグローバル変数に設定（テスト用）
	db.DB = mockDB

	// モックデータ
	mockTodo := domain.Todo{Name: "Test Todo"}
	mockResponse := domain.Todo{ID: 1, Name: "Test Todo"}

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
	if assert.NoError(t, HandleAddTodo(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディの検証
		var responseTodo domain.Todo
		err := json.Unmarshal(rec.Body.Bytes(), &responseTodo)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, responseTodo)
	}

	// モックの期待値の検証
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandleDeleteTodo(t *testing.T) {
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
		err := HandleDeleteTodo(c)
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
		err := HandleDeleteTodo(c)
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
		err := HandleDeleteTodo(c)
		assert.NoError(t, err)

		// ステータスコードとエラーメッセージが期待通り
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "データベースエラー", rec.Body.String())

		// モックDBの期待通りに動作したかを検証
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
