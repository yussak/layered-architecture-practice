package domain_test

import (
	"errors"
	"reflect"
	"server/domain"
	"server/infrastructure"
	"testing"
)

// MockTodoRepo: infrastructure.TodoRepo を満たすモック
type MockTodoRepo struct {
	Todos []infrastructure.Todo
	Err   error
}

func (m *MockTodoRepo) GetTodosFromDB() ([]infrastructure.Todo, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Todos, nil
}

func TestGetTodos(t *testing.T) {
	// モックデータを定義
	mockInfraTodos := []infrastructure.Todo{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
	}

	mockRepo := &MockTodoRepo{
		Todos: mockInfraTodos,
	}

	repo := &domain.Repo{Repo: mockRepo}

	t.Run("Success", func(t *testing.T) {
		todos, err := domain.GetTodos(repo)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// domain.Todo に変換した期待値を用意
		expectedTodos := []domain.Todo{
			{ID: 1, Name: "Task 1"},
			{ID: 2, Name: "Task 2"},
		}

		if !reflect.DeepEqual(todos, expectedTodos) {
			t.Errorf("expected %v, got %v", expectedTodos, todos)
		}
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.Err = errors.New("mock error")
		_, err := domain.GetTodos(repo)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
		if err.Error() != "mock error" {
			t.Errorf("expected 'mock error', got %v", err)
		}
	})
}
