package application

import (
	"server/infrastructure"
	"server/models"
)

func GetTodos() ([]models.Todo, error) {
	todos, err := infrastructure.GetTodosFromDB()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
