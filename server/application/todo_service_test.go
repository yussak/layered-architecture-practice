package application

import (
	"errors"
	"reflect"
	"server/domain"
	"testing"
)

type MockTodoRepository struct {
	Called bool
	Todos  []domain.Todo
	Error  error
}

func (m *MockTodoRepository) GetTodos() ([]domain.Todo, error) {
	m.Called = true
	return m.Todos, m.Error
}

func TestTodoServiceImpl_GetTodos_Success(t *testing.T) {
	// モックリポジトリを作成
	mockRepo := &MockTodoRepository{
		Todos: []domain.Todo{
			{ID: 1, Name: "Test Todo 1"},
			{ID: 2, Name: "Test Todo 2"},
		},
		Error: nil,
	}

	// サービスにモックを注入
	TodoService := TodoService(mockRepo)

	// 実行
	result, err := TodoService.GetTodos()

	// 検証
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, mockRepo.Todos) {
		t.Errorf("expected %v, got %v", mockRepo.Todos, result)
	}

	if !mockRepo.Called {
		t.Error("expected GetTodos to be called, but it was not")
	}
}

func TestTodoServiceImpl_GetTodos_Error(t *testing.T) {
	mockError := errors.New("mock error")
	// モックリポジトリを作成
	mockRepo := &MockTodoRepository{
		Todos: nil,
		Error: mockError,
	}

	// サービスにモックを注入
	TodoService := TodoService(mockRepo)

	// 実行
	result, err := TodoService.GetTodos()

	if err == nil {
		t.Fatal("expected error, but got nil")
	}

	// 検証
	if err != mockError {
		t.Errorf("unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}

	if !mockRepo.Called {
		t.Errorf("expected GetTodos to be called, but it was not")
	}
}
