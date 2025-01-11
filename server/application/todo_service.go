package application

import (
	"server/domain"
)

func GetTodos() ([]domain.Todo, error) {
	return domain.GetTodos()
}

func CreateTodo(name string) (domain.Todo, error) {
	return domain.GetNewTodo(name)
}
