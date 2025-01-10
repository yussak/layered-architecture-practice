package application

import (
	"server/domain"
)

func GetTodos() ([]domain.Todo, error) {
	return domain.GetTodos()
}
