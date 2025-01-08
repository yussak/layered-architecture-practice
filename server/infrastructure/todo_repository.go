package infrastructure

import (
	"server/db"
	"server/models"
)

func GetTodosFromDB() ([]models.Todo, error) {
	rows, err := db.DB.Query("SELECT id, name FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		todo := models.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Name); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
