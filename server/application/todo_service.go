package application

// ユースケースを書く

import (
	"server/domain"
)

func GetTodos() ([]domain.Todo, error) {
	return domain.GetTodos()
}

func CreateTodo(name string) (domain.Todo, error) {
	return domain.GetNewTodo(name)
}

func DeleteTodo(id string) error {
	return domain.DeleteTodo(id)
}
